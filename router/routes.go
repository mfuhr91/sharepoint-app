package router

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"sharepoint-app/controllers"
	"sharepoint-app/repositories"
	"sharepoint-app/services"
	"sharepoint-app/templates"
	"strconv"
)

var (
	filesRepository = repositories.NewFilesRepository()
	filesService    = services.NewFilesService(filesRepository)
	filesController = controllers.NewFilesController(filesService)
)

func InitRoutes(r *gin.Engine) {
	r.NoRoute(templates.NoRouteHandler)
	r.GET("/ping", templates.Ping)
	
	getMB := template.FuncMap{
		"getMB": func(bytes int64) string {
			res := bytes / 1024 / 1024
			if res == 0 {
				return strconv.FormatInt(bytes/1024, 10) + " KB"
			}
			return strconv.FormatInt(res, 10) + " MB"
		},
	}
	r.SetFuncMap(getMB)
	r.Static("assets", "resources/assets")
	
	r.LoadHTMLGlob("resources/templates/*")
	
	r.GET("/", templates.FilesRender)
	
	r.POST("/files/save", filesController.Save)
	r.POST("/files/delete/:id", filesController.Delete)
	
}
