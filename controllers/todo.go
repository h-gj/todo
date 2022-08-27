package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"todo/services"
)

func CreateTodo(c *gin.Context) {
	var service = services.CreateTodoService{}
	err := c.ShouldBindJSON(&service)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1004,
			"msg":  err.Error(),
		})
		return
	}
	if err := service.Create(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1000,
			"msg":  "Internal server error.",
		})
	}
	//err := c.ShouldBind(&createTodoService)
	//fmt.Println(23232, err)
	c.JSON(200, gin.H{
		"msg": "success",
	})
}
