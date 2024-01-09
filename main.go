package main

import (
	"github.com/Kachyr/crud/controllers"
	"github.com/Kachyr/crud/initializers"
	"github.com/Kachyr/crud/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	// Authentication
	r.POST("/singup", controllers.SingUp)
	r.POST("/login", controllers.LogIn)
	// Posts
	r.POST("/posts", middleware.RequireAuth, controllers.PostsCreate)
	r.GET("/posts", controllers.GetAllPosts)
	r.GET("/posts/:id", controllers.GetPost)
	r.PUT("/posts/:id", middleware.RequireAuth, controllers.UpdatePost)
	r.DELETE("/posts/:id", middleware.RequireAuth, controllers.DeletePost)
	r.Run() // listen and serve on 0.0.0.0:ENV.PORT
}
