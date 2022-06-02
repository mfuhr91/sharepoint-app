package models

import "time"

type File struct {
	ID          string    `bson:"_id"`
	Name        string    `bson:"Name"`
	Ext         string    `bson:"Ext"`
	Size        int64     `bson:"Size"`
	Icon        string    `bson:"Icon"`
	ViewUrl     string    `bson:"ViewUrl"`
	DownloadUrl string    `bson:"DownloadUrl"`
	PublicID    string    `bson:"PublicID"`
	Time        time.Time `bson:"Time"`
}
