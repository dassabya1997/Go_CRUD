package main

import (
	"go_crud/initializers"
	model "go_crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&model.Post{})
}
