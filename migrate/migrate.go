package main

import (
	"home/dsi/go/bin/go-prac/initializers"
	"home/dsi/go/bin/go-prac/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
