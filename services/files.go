package services

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"log"
	"mime/multipart"
	"os"
	"sharepoint-app/models"
	"sharepoint-app/repositories"
	"strings"
	"time"
)

type fileService struct{}

type FileService interface {
	GetAll() ([]models.File, error)
	GetById(id string) (models.File, error)
	Save(*multipart.FileHeader) (models.File, error)
	Delete(id string) error
}

var (
	fileRepository repositories.FileRepository
	
	cloudName = os.Getenv("CLOUDINARY_CLOUD_NAME")
	apikey    = os.Getenv("CLOUDINARY_API_KEY")
	apiSecret = os.Getenv("CLOUDINARY_API_SECRET")
)

func NewFilesService(repository repositories.FileRepository) FileService {
	fileRepository = repository
	return &fileService{}
}

func (c fileService) GetAll() ([]models.File, error) {
	files, err := fileRepository.GetAll()
	if err != nil {
		log.Printf("error when getting the files from service")
		return files, err
	}
	
	loc, _ := time.LoadLocation("America/Argentina/Buenos_Aires")
	for i, file := range files {
		files[i].Time = file.Time.In(loc)
	}
	
	return files, nil
}

func (c fileService) Save(formHeader *multipart.FileHeader) (models.File, error) {
	var err error
	var file models.File
	
	fileNameSlice := strings.Split(formHeader.Filename, ".")
	
	file.Name = strings.Join(fileNameSlice[:len(fileNameSlice)-1], ".")
	file.Ext = fileNameSlice[len(fileNameSlice)-1]
	file.Size = formHeader.Size
	file.Icon = getIcon(formHeader)
	file.Time = time.Now()
	
	formFile, err := formHeader.Open()
	if err != nil {
		return models.File{}, err
	}
	uploadResult, err := UploadFile(formFile, file)
	if err != nil {
		return models.File{}, err
	}
	
	urlSplited := strings.Split(uploadResult.SecureURL, "upload")
	downloadUrl := urlSplited[0] + "upload/fl_attachment" + urlSplited[1]
	
	file.PublicID = uploadResult.PublicID
	file.DownloadUrl = downloadUrl
	file.ViewUrl = uploadResult.SecureURL
	file, err = fileRepository.Save(file)
	
	if err != nil {
		return models.File{}, err
	}
	
	return file, nil
}

func (c *fileService) GetById(id string) (models.File, error) {
	file, err := fileRepository.GetById(id)
	if err != nil {
		return models.File{}, err
	}
	
	return file, nil
}

func (c fileService) Delete(id string) error {
	file, err := fileRepository.GetById(id)
	if err != nil {
		return err
	}
	
	cld, err := cloudinary.NewFromParams(cloudName, apikey, apiSecret)
	if err != nil {
		return err
	}
	
	err = fileRepository.Delete(id)
	if err != nil {
		return err
	}
	
	uploadResult, err := cld.Upload.Destroy(context.TODO(), uploader.DestroyParams{PublicID: file.PublicID})
	if err != nil {
		return err
	}
	
	if uploadResult.Result == "ok" {
		log.Printf("file deleted successfully from cloudinary - result: %+v", uploadResult)
	} else {
		log.Printf("error file has not deleted from cloudinary - result: %+v", uploadResult)
	}
	
	return nil
}

func UploadFile(file multipart.File, modelFile models.File) (*uploader.UploadResult, error) {
	
	cld, err := cloudinary.NewFromParams(cloudName, apikey, apiSecret)
	if err != nil {
		return &uploader.UploadResult{}, err
	}
	
	uploadResult, err := cld.Upload.Upload(context.TODO(), file, uploader.UploadParams{Folder: "sharepoint-app", ResourceType: "auto", PublicID: modelFile.Name})
	if err != nil {
		return &uploader.UploadResult{}, err
	}
	
	log.Printf("file uploaded successfully to cloudinary - result: %+v", uploadResult)
	
	return uploadResult, nil
	
}

func getIcon(fileName *multipart.FileHeader) string {
	
	contentType := fileName.Header.Get("Content-Type")
	
	if strings.Contains(contentType, "pdf") {
		return "fa-file-pdf"
	}
	if strings.Contains(contentType, "audio") {
		return "fa-file-audio"
	}
	if strings.Contains(contentType, "image") {
		return "fa-file-image"
	}
	if strings.Contains(contentType, "video") {
		return "fa-file-video"
	}
	if strings.Contains(contentType, "powerpoint") || strings.Contains(contentType, "presentation") {
		return "fa-file-powerpoint"
	}
	if strings.Contains(contentType, "excel") || strings.Contains(contentType, "spreadsheet") {
		return "fa-file-excel"
	}
	if strings.Contains(contentType, "word") || strings.Contains(contentType, "opendocument.text") {
		return "fa-file-word"
	}
	if strings.Contains(contentType, "compressed") ||
		strings.Contains(contentType, "zip") ||
		strings.Contains(contentType, "tar") {
		return "fa-file-zipper"
	}
	if strings.Contains(contentType, "html") ||
		strings.Contains(contentType, "xml") {
		return "fa-file-code"
	}
	
	return "fa-file"
	
}
