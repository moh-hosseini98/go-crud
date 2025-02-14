package models
import "gorm.io/gorm"

type Todo structs {
	gorm.Model
	Content string
	Status bool
}