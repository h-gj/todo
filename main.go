package main


import (
	"todo/routes"
)


type CreateToDoService struct {
	Content string `form:"content" binding:"required"`
}

func main() {
	r := routes.InitRoutes()
	db := InitDb()
	r.Run()
}


