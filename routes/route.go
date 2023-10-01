package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/Ilmanzi/task-5-pbi-btpns-ilmanaz.git/controllers"
	"github.com/Ilmanzi/task-5-pbi-btpns-ilmanaz.git/middleware"
)

func Routes() *gin.Engine {

	router := gin.Default()

	app := router.Group("/api/v1")

	app.POST("/auth/login", controllers.Login)
	app.POST("/auth/register", controllers.Register)
	app.Static("/pictures", "./uploads")
	app.Use(middleware.AuthMiddleware())
	app.PUT("/users/:id", controllers.UpdateUser)
	app.DELETE("/users/:id", controllers.DeleteUser)
	app.GET("/pictures", controllers.GetPicture)
	app.POST("/pictures", controllers.UploadPicture)
	app.PUT("/pictures/:id", controllers.UpdatePicture)
	app.DELETE("/pictures/:id", controllers.DeletePicture)

	return router
}
