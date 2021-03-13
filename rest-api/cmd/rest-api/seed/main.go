package main

import (
	"example.com/rest-api/internal/apps/models"
)

func main() {
	models.Connect()
	models.Seed()
}
