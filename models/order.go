package models

type Order struct {
	Name           string       `json:"name"`
	Phone          string       `json:"phone"`
	ShoppingCartID int          `json:"id_shoppingcart"`
	ShoppingCart   ShoppingCart `gorm:"foreignKey:ShoppingCartID"`
}
