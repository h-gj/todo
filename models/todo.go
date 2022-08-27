package models

import (
	"log"
)

type ToDo struct {
	//gorm.Model
	Content string
	F       string
}

func CreateTodo(content string, f string) error {
	var todo ToDo
	todo = ToDo{Content: content, F: f}
	res := db.Model(&ToDo{}).Create(&todo)
	if res.Error != nil {
		log.Fatal(res.Error)
		return res.Error
	}

	return nil
}
