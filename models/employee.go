package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name       string `json:"name" gorm:"not null"`
	Position   string `json:"position" gorm:"not null"`
	Department string `json:"department" gorm:"not null"`
	Salary     int    `json:"salary" gorm:"not null"`
	Email      string `json:"email" gorm:"not null;unique"`
	Phone      string `json:"phone" gorm:"not null;unique"`
	Address    string `json:"address" gorm:"not null"`
}
