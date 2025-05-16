package routes

import (
	"github.com/gin-gonic/gin"
	controllers "go-mongo-crud/controller"
)

func UserRoute(router *gin.Engine) {
	user := router.Group("/user")
	{
		user.POST("/", controllers.CreateUser)
		user.GET("/", controllers.GetUsers)
		user.PUT("/:id", controllers.UpdateUser)
		user.DELETE("/:id", controllers.DeleteUser)
	}
}