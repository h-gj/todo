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

		c.JSON(http.StatusBadRequest, nil)
		return
	}
	//err := c.ShouldBind(&createTodoService)
	//fmt.Println(23232, err)
	c.JSON(200, gin.H{
		"msg": "success",
	})
}