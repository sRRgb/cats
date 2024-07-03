package models

import (
	"github.com/jinzhu/gorm"
)

type Cat struct {
	gorm.Model
	Name              string  `json:"name"`
	YearsOfExperience int     `json:"years_of_experience"`
	Breed             string  `json:"breed"`
	Salary            float64 `json:"salary"`
}

type CreateCatInput struct {
	Name              string  `json:"name" binding:"required"`
	YearsOfExperience int     `json:"years_of_experience" binding:"required"`
	Breed             string  `json:"breed" binding:"required"`
	Salary            float64 `json:"salary" binding:"required"`
}

type UpdateCatInput struct {
	Salary float64 `json:"salary" binding:"required"`
}
