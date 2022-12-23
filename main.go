package main

import (
	"home/dsi/go/bin/go-prac/controllers"
	"home/dsi/go/bin/go-prac/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostCreate)
	r.GET("/all", controllers.PostIndex)
	r.GET("/get/:id", controllers.PostShow)
	r.PUT("updatePost/:id", controllers.PostUpdate)
	r.DELETE("deletePost/:id", controllers.PostDelete)
	r.Run()
}
