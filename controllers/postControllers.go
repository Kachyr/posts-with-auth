package controllers

import (
	"fmt"

	"github.com/Kachyr/crud/initializers"
	"github.com/Kachyr/crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// Create post
	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		fmt.Println("Error creating post:", result.Error)
		return
	}

	// res

	c.JSON(200, post)
}

func GetAllPosts(c *gin.Context) {
	posts := []models.Post{}
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.Status(400)
		fmt.Println("Error cant find posts:", result.Error)
		return
	}

	c.JSON(200, posts)
}

func GetPost(c *gin.Context) {
	id := c.Param("id")
	post := models.Post{}
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.Status(400)
		fmt.Println("Error cant find posts:", result.Error)
		return
	}

	c.JSON(200, post)
}

func UpdatePost(c *gin.Context) {
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	id := c.Param("id")
	post := models.Post{}
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.Status(400)
		fmt.Println("Error cant find posts:", result.Error)
		return
	}

	post.Title = body.Title
	post.Body = body.Body

	result = initializers.DB.Save(&post)
	if result.Error != nil {
		c.Status(400)
		fmt.Println("Error cant update posts:", result.Error)
		return
	}

	c.JSON(200, post)
}

func DeletePost(c *gin.Context) {

	id := c.Param("id")
	result := initializers.DB.Delete(&models.Post{}, id)

	if result.Error != nil {
		c.Status(400)
		fmt.Println("Error cant delete post:", result.Error)
		return
	}

	c.String(200, "Success")
}
