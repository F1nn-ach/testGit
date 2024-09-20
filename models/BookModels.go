package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	BookName string  `gorm:"not null"`
	UrlImage string  `gorm:"not null"`
	Price    float64 `gorm:"not null"`
}
