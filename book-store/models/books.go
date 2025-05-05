package models

type Book struct {
	Id     string `gorm:"primaryKey"`
	Title  string
	Author string
	Price  float64
}
