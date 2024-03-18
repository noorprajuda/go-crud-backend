package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/noorprajuda/go-crud/initializers"
	"github.com/noorprajuda/go-crud/models"
)

func PostsCreate(c *gin.Context) {

	// Get data of req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Create a post

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return it

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with them

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Get the posts
	var post models.Post
	initializers.DB.Find(&post, id)

	// Respond with them

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Get the data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Find the post were updating
	var post models.Post
	initializers.DB.Find(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// Respond with them

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Delete the post
	initializers.DB.Delete(&models.Post{}, id)

	// Respond with them
	c.JSON(200, gin.H{
		"message": `Post with id ` + id + ` has been deleted`,
	})
}
