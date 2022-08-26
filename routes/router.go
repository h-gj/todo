package routes

import (
	"github.com/gin-gonic/gin"
	"todo/controllers"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("", controllers.CreateTodo)
	return r
}
