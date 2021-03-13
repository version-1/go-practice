package main

import (
	route "example.com/rest-api/internal/apps/routes"
	"example.com/rest-api/internal/apps/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.Connect()
	route.Register(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func setup() {

}
