package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	ProductID     uint
	CarpetColorID int         `json:"id_carpetcolor"`
	CarpetColor   CarpetColor ` gorm:"foreignKey:CarpetColorID"`
	Size          string
	Quantity      uint
	Price         uint64
	OrderID       int   `json:"id_order"`
	Order         Order ` gorm:"foreignKey:OrderID"`
}
