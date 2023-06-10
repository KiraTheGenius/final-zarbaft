package models

import "gorm.io/gorm"

type Carpet struct {
	gorm.Model
	NameFa       string     `gorm:"not null" db:"namefa" json:"nameFa"`
	NameEn       string     `gorm:"not null" db:"namefa" json:"nameEn"`
	Shane        float64    `gorm:"not null" db:"shane" json:"shane"`
	Density      string     `gorm:"not null" db:"density" json:"density"` // Tarakom
	StyleFa      string     `gorm:"not null" db:"style" json:"styleFa"`   // Noe
	StyleEn      string     `gorm:"not null" db:"style" json:"styleEn"`   // Noe
	CodeNaqshe   string     `gorm:"not null" db:"code_naqshe" json:"code_naqshe"`
	CollectionID int64      `gorm:"not null"`
	Collection   Collection `gorm:"foreignkey:CollectionID;references:ID"`
	Slug         string     `gorm:"not null" db:"slug" json:"slug"`
	MostPopular  float64    `gorm:"not null"`
	BestSelling  float64    `gorm:"not null"`
}
