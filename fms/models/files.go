package models

type Folders struct {
	Id     string `gorm:"primaryKey" json:"id"`
	Name   string `json:"name"`
	UserId string `json:"user_id"`
}

type Files struct {
	Id       string `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	FolderId string `json:"folder_id"`
}
