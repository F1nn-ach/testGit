package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	BookId   int    `goam:"unique"`
	BookName string `goam:"not null"`
	UrlImage string `goam:"not null"`
	Price    int    `goam:"not null"`
}

type ListBook struct {
}
