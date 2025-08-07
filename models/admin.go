package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"` // ซ่อน password ไม่ให้แสดงเวลา return JSON
	Role     string `json:"role" gorm:"type:varchar(20);default:'admin'"`
}
