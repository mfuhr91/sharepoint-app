package templates

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sharepoint-app/repositories"
	"sharepoint-app/services"
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

/*func DeleteConfirmRender(c *gin.Context) {
	id := c.Query("id")

	c.HTML(http.StatusOK, "delete-confirm.html", gin.H{
		"Title": "Eliminar Archivo",
		"ID":    id,
	})
	return

}*/
