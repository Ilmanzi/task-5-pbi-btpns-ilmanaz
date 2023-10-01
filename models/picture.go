package models

import (
	"gorm.io/gorm"
)

type Picture struct {
	gorm.Model
	Title      string `json:"title"`
	Caption    string `json:"caption"`
	PictureUrl string `json:"picture_url"`
	UserID     string `json:"user_id"`
	User       User   `gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}
