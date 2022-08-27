package services

import "todo/models"

type CreateTodoService struct {
	Content string `json:"content" binding:"required"`
	F       string
}

func (s *CreateTodoService) Create() error {
	return models.CreateTodo(s.Content, s.F)
}
