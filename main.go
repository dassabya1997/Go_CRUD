package main

import (
	// "log"

	"go_crud/controllers"
	"go_crud/initializers"

	"github.com/gin-gonic/gin"
	//"github.com/joho/godotenv"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()

	// before controller

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome!!",
		})
	})

	// after controllers
	r.POST("/createPost", controllers.PostsCreate)
	r.GET("/getPosts", controllers.GetPosts)
	r.GET("/getPosts/:id", controllers.GetSinglePost)
	r.DELETE("/deletePost/:id", controllers.DeletePost)
	r.PUT("/UpdatePost/:id", controllers.UpdatePost)
	r.Run() // listen and serve on 0.0.0.0:8080
}
