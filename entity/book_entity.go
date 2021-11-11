package entity

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string `gorm:"column:title"`
	Author string `gorm:"column:author"`
	Year   int    `gorm:"column:year"`
}

func (Book) TableName() string {
	return "books"
}
