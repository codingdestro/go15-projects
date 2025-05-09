package models

type Users struct {
	Id       string `gorm:"PrimaryKey" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
