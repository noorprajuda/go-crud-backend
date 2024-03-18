package main

import (
	"github.com/noorprajuda/go-crud/initializers"
	"github.com/noorprajuda/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
