package models

import . "todo/services"

type ToDo struct{
	Content string
}


func CreateTodo(s *CreateTodoService) *ToDo {

	//ToDo{
	//	Content: s.Content,
	//}
}