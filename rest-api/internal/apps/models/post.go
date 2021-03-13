package models

import (
	"gorm.io/gorm"
)

type Post struct {
  gorm.Model
  UserId uint
  Title string
  Content string
}
