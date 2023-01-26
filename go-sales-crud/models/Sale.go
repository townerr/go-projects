package models

type Sale struct {
	Id     int    `json:"id" gorm:"primary_key"`
	UserId int    `json:"user"`
	Amount string `json:"amount"`
}
