package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"starbucks-app/services"
	"starbucks-app/utils"
)

type fileController struct{}

type FileController interface {
	Save(c *gin.Context)
	Delete(c *gin.Context)
}

var (
	fileService services.FileService
)

func NewFilesController(service services.FileService) FileController {
	fileService = service
	return &fileController{}
}

func (cont *fileController) Save(c *gin.Context) {
	
	formFile, err := c.FormFile("NewFile")
	if err != nil {
		return
	}
	
	_, err = fileService.Save(formFile)
	if err != nil {
		log.Printf("error when saving file - %s", err)
		return
	}
	
	utils.Redirect("/files", c)
}

func (cont *fileController) Delete(c *gin.Context) {
	id := c.Query("id")
	err := fileService.Delete(id)
	if err != nil {
		log.Printf("error when deleting file - %s", err)
		return
	}
	
	utils.Redirect("/files", c)
}
