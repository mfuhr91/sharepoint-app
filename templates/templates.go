package templates

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"starbucks-app/repositories"
	"starbucks-app/services"
)

var (
	filesRepository = repositories.NewFilesRepository()
	filesService    = services.NewFilesService(filesRepository)
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}

func NoRouteHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "error.html", gin.H{
		"Error": "Esta página no existe!",
		"Msg":   "Toca el botón para reintentar",
	})
	return
}

func FilesRender(c *gin.Context) {
	files, err := filesService.GetAll()
	
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"Error": err,
			"Msg":   "Toca el botón para reintentar",
		})
		return
	}
	
	c.HTML(http.StatusOK, "files.html", gin.H{
		"Title": "Archivos",
		"Files": files,
	})
}

func FileRenderUpload(c *gin.Context) {
	
	c.HTML(http.StatusOK, "files-upload.html", gin.H{
		"Title": "Subir Archivo",
	})
	return
}

func FileRenderView(c *gin.Context) {
	id := c.Param("id")
	file, err := filesService.GetById(id)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"Error": err,
			"Msg":   "Toca el botón para reintentar",
		})
		return
	}
	
	c.HTML(http.StatusOK, "files-view.html", gin.H{
		"Title":    "Subir Archivo",
		"FilePath": file.ViewUrl,
		"File":     file,
	})
	return
}

func DeleteConfirmRender(c *gin.Context) {
	id := c.Query("id")
	
	c.HTML(http.StatusOK, "delete-confirm.html", gin.H{
		"Title": "Eliminar Archivo",
		"ID":    id,
	})
	return
	
}
