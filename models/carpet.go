package models

import "gorm.io/gorm"

type Carpet struct {
	gorm.Model   `json:"-"`
	NameFa       string  `gorm:"not null; column:name_fa" db:"namefa" json:"nameFa" `
	NameEn       string  `gorm:"not null" db:"namefa" json:"nameEn,omitempty"`
	Shane        float64 `gorm:"not null; column:shane" db:"shane" json:"shane,omitempty"`
	Density      string  `gorm:"not null; column:density" db:"density" json:"density,omitempty"` // Tarakom
	StyleFa      string  `gorm:"not null" db:"style" json:"styleFa,omitempty"`                   // Noe
	StyleEn      string  `gorm:"not null" db:"style" json:"styleEn,omitempty"`                   // Noe
	CodeNaqshe   string  `gorm:"not null" db:"code_naqshe" json:"code_naqshe,omitempty"`
	CollectionID int64   `gorm:"not null"`
	Slug         string  `gorm:"not null" db:"slug" json:"slug,omitempty"`
	MostPopular  float64 `gorm:"not null" json:"most_popular,omitempty"`
	BestSelling  float64 `gorm:"not null" json:"best_selling,omitempty"`
}
