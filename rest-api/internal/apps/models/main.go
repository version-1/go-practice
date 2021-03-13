package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
  dsn := "root:@tcp(127.0.0.1:3306)/go-practice-rest-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
  if err != nil {
    panic("failed to connect database")
  }

	var users []User
	db.Find(&users)
	if len(users) > 0 {
	  DB = db
		return
	}

	DB = db
}

func Migrate() {
	DB.AutoMigrate(&User{}, &Category{}, &Post{}, &PostCategory{})
}

func Seed() {
	// Create
	email := "jhon@example.com"
	DB.Create(&User{Name: "Jhon", Email: &email })
	var user User
	DB.First(&user, "email = ?", email)

	description := ""
  DB.Create(&Category{Name: "Programming", Description: &description })
  DB.Create(&Category{Name: "Column", Description: &description })

	DB.Create(&Post{UserId: user.ID, Title: "Getting Started", Content: "Lorem ipsum" })
	DB.Create(&Post{UserId: user.ID, Title: "Getting Started", Content: "Lorem ipsum" })
}
