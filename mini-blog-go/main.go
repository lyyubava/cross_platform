package main

import (
	"mini-blog-go/mini-blog-go/controllers"
	"mini-blog-go/mini-blog-go/controllers/auth"
	"mini-blog-go/mini-blog-go/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4000"} // Replace with your frontend URL
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	r.Use(cors.New(config))
	router_posts := r.Group("/post")
	//router_posts.Use(middlewares.JwtAuthMiddleware())
	//router_posts_pub.GET("/", controllers.FindPosts)
	//router_posts_private := .Group("/posts")
	router_posts.POST("/create", controllers.CreatePost)
	router_posts.PUT("/edit", controllers.EditPost)
	router_posts.DELETE("/delete", controllers.DeletePost)
	router_posts.GET("/", controllers.GetPost)
	router_posts.GET("/all", controllers.GetAllPosts)

	router_auth := r.Group("/auth")
	router_auth.POST("/register", auth.Register)
	router_auth.POST("/login", auth.Login)
	r.Run("localhost:8080")

}
