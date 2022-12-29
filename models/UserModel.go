package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name       string
	RollNumber int
	Email      string
}
