package repositories

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sharepoint-app/database"
	"sharepoint-app/models"
)

type fileRepository struct{}

type FileRepository interface {
	GetAll() ([]models.File, error)
	GetById(id string) (models.File, error)
	Save(models.File) (models.File, error)
	Delete(id string) error
}

const (
	databaseName   = "sharepointdb"
	collectionName = "files"
)

func NewFilesRepository() FileRepository {
	return &fileRepository{}
}

func (f *fileRepository) GetAll() ([]models.File, error) {
	client, err := database.MongoConnect()
	if err != nil {
		return []models.File{}, err
	}
	
	collection := client.Database(databaseName).Collection(collectionName)
	
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return []models.File{}, err
	}
	var files []models.File
	
	err = cursor.All(context.TODO(), &files)
	if err != nil {
		return []models.File{}, err
	}
	
	return files, nil
}

func (f *fileRepository) Save(file models.File) (models.File, error) {
	client, err := database.MongoConnect()
	if err != nil {
		return models.File{}, err
	}
	
	collection := client.Database(databaseName).Collection(collectionName)
	
	_, err = collection.InsertOne(context.TODO(),
		bson.M{
			"Name":        file.Name,
			"Ext":         file.Ext,
			"Size":        file.Size,
			"Icon":        file.Icon,
			"ViewUrl":     file.ViewUrl,
			"DownloadUrl": file.DownloadUrl,
			"PublicID":    file.PublicID,
			"Time":        file.Time,
		})
	if err != nil {
		return models.File{}, err
	}
	
	return file, nil
}

func (f *fileRepository) GetById(id string) (models.File, error) {
	var file models.File
	
	client, err := database.MongoConnect()
	if err != nil {
		return models.File{}, err
	}
	
	collection := client.Database(databaseName).Collection(collectionName)
	objID, _ := primitive.ObjectIDFromHex(id)
	err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&file)
	if err != nil {
		return models.File{}, err
	}
	
	return file, nil
}

func (f *fileRepository) Delete(id string) error {
	client, err := database.MongoConnect()
	if err != nil {
		return err
	}
	
	collection := client.Database(databaseName).Collection(collectionName)
	objID, _ := primitive.ObjectIDFromHex(id)
	
	res, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		return err
	}
	
	if res.DeletedCount == 0 {
		return errors.New(fmt.Sprintf("document id: %s not found", id))
	}
	
	return nil
}
