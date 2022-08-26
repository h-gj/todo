package services

import "todo/models"

type CreateTodoService struct{
	Content string
}


func(s *CreateTodoService) Create() {
	models.CreateTodo(s)
}