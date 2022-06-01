package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
		c.HTML(http.StatusOK, "error.html", gin.H{
			"Error": err,
			"Msg":   "Toca el botón para reintentar",
		})
		return
	}
	
	_, err = fileService.Save(formFile)
	if err != nil {
		log.Printf("error when saving the file - %s", err)
		return
	}
	
	utils.Redirect("/files", c)
}

func (cont *fileController) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"Error": errors.New("debe enviar un ID"),
			"Msg":   "Toca el botón para reintentar",
		})
	}
	
	err := fileService.Delete(id)
	if err != nil {
		log.Printf("error when deleting the file - %s", err)
		return
	}
	
	utils.Redirect("/files", c)
}
