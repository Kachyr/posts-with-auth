package controllers

import (
	"strconv"

	"github.com/Kachyr/crud/helpers"
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
		println("Error creating post:", result.Error)
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
		println("Error cant find posts:", result.Error)
		return
	}

	c.JSON(200, posts)
}

func GetPostsPaginated(c *gin.Context) {
	posts := []models.Post{}
	println(posts)
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.Status(400)
		println("Error: ", err)
		return
	}

	pageSizeStr := c.DefaultQuery("size", "10")
	size, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		c.Status(400)
		println("Error: ", err)
		return
	}

	offset := (page - 1) * size

	result := initializers.DB.Limit(size).Offset(offset).Find(&posts)

	if result.Error != nil {
		c.Status(400)
		println("Error cant find posts:", result.Error)
		return
	}

	var totalElements int64
	initializers.DB.Model(&models.Post{}).Count(&totalElements)

	totalPages := helpers.CalculateTotalPages(int(totalElements), size)

	c.JSON(200, models.PaginatedContent[models.Post]{
		Data:       posts,
		Page:       page,
		TotalPages: totalPages,
	})
}

func GetPost(c *gin.Context) {
	id := c.Param("id")
	post := models.Post{}
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.Status(400)
		println("Error cant find posts:", result.Error)
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
		println("Error cant find posts:", result.Error)
		return
	}

	post.Title = body.Title
	post.Body = body.Body

	result = initializers.DB.Save(&post)
	if result.Error != nil {
		c.Status(400)
		println("Error cant update posts:", result.Error)
		return
	}

	c.JSON(200, post)
}

func DeletePost(c *gin.Context) {

	id := c.Param("id")
	result := initializers.DB.Delete(&models.Post{}, id)

	if result.Error != nil {
		c.Status(400)
		println("Error cant delete post:", result.Error)
		return
	}

	c.String(200, "Success")
}
