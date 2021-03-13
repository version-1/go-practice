package routes

import (
	"errors"

	"example.com/rest-api/internal/apps/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register (r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	v1 := r.Group("/api/v1")
	groups := [...]func(r *gin.RouterGroup){posts, user, categories}

	for _, g := range groups {
		g(v1)
	}

  return
}

func categories (r *gin.RouterGroup) {
	g := r.Group("/categories")
	g.GET("/", func(c *gin.Context) {
		var categories []models.Category
	  models.DB.Find(&categories)
		c.JSON(200, gin.H{
			"data": categories,
		})
	})
	g.GET("/:id", func(c *gin.Context) {
		var category []models.Category
		id := c.Param("id")
		result := models.DB.First(&category, id)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{})
			return
		}
		c.JSON(200, gin.H{
			"data": category,
		})
	})
	g.POST("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	g.PATCH("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	g.DELETE("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func user (r *gin.RouterGroup) {
	g := r.Group("/user")
	g.GET("/", func(c *gin.Context) {
		var user models.User
		models.DB.First(&user)
		c.JSON(200, gin.H{
			"data": user,
		})
	})
	g.PATCH("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func posts (r *gin.RouterGroup) {
	g := r.Group("/posts")

	g.GET("/", func(c *gin.Context) {
		var posts []models.Post
	  models.DB.Find(&posts)
		c.JSON(200, gin.H{
			"data": posts,
		})
	})
	g.GET("/:id", func(c *gin.Context) {
		var post models.Post
		id := c.Param("id")
		result := models.DB.First(&post, id)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{})
			return
		}
		c.JSON(200, gin.H{
			"data": post,
		})
	})
	g.POST("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	g.PATCH("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	g.DELETE("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
