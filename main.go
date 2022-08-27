package main

import (
	m "todo/models"
	"todo/pkg/settings"
	"todo/routes"
)

func main() {
	r := routes.InitRoutes()
	settings.SetUp()
	m.SetUp()
	InitDb()
	r.Run()
}
