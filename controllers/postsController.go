package controllers

import (
	"fmt"
	"go_crud/initializers"
	model "go_crud/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// get data off req body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)
	// create a post
	post := model.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		// If there's an error, print it and return a 400 status
		fmt.Println("Error creating post:", result.Error)
		c.Status(400)
		return
	}

	// return the post
	c.JSON(200, gin.H{"post": post})
}

func GetPosts(c *gin.Context) {
	//get all posts
	var posts []model.Post
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		// If there's an error, print it and return a 400 status
		fmt.Println("Error fetching all posts:", result.Error)
		c.Status(400)
		return
	}

	// return the post
	c.JSON(200, gin.H{"post": posts})
}

func GetSinglePost(c *gin.Context) {
	// get the postId
	id := c.Param("id")

	var post model.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		// If there's an error, print it and return a 400 status
		fmt.Println("Error fetching a single post:", result.Error)
		c.Status(400)
		return
	}

	// return the post
	c.JSON(200, gin.H{"post": post})
}

func DeletePost(c *gin.Context) {
	// get the postId
	id := c.Param("id")

	var post model.Post
	result := initializers.DB.Delete(&post, id)

	if result.Error != nil {
		// If there's an error, print it and return a 400 status
		fmt.Println("Error fetching a single post:", result.Error)
		c.Status(400)
		return
	}

	// return the post
	c.JSON(200, gin.H{"message": "Post deleted."})
}

func UpdatePost(c *gin.Context) {
	// get the postId
	id := c.Param("id")

	// get the data off req body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)
	// create a post
	var post model.Post

	// find with which record you want to update
	initializers.DB.First(&post, id)
	// update it
	initializers.DB.Model(&post).Updates(model.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// return the post
	c.JSON(200, gin.H{"post": post})
}
