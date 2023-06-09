package models

import "github.com/jinzhu/gorm"

type ShoppingCart struct {
	gorm.Model
	ID    int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Token string `json:"token"`
	// Items []CartItem
}
