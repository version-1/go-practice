package models

import (
	"gorm.io/gorm"
)

type PostCategory struct {
  gorm.Model

  PostId string
  CategoryId *string
}
