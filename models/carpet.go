package models

import "gorm.io/gorm"

type Carpet struct {
	gorm.Model   `json:"-" fa:"-" en:"-"`
	NameFa       string  `gorm:"not null; column:name_fa" db:"namefa" json:"nameFa" fa:"name,omitempty" en:"trash_name,omitempty" `
	NameEn       string  `gorm:"not null" db:"namefa" json:"nameEn,omitempty" fa:"trash_name,omitempty" en:"name,omitempty"`
	Shane        float64 `gorm:"not null; column:shane" db:"shane" json:"shane,omitempty" fa:"shane,omitempty" en:"shane,omitempty"`
	Density      string  `gorm:"not null; column:density" db:"density" json:"density,omitempty" fa:"density,omitempty" en:"density,omitempty"` // Tarakom
	StyleFa      string  `gorm:"not null" db:"style" json:"styleFa,omitempty" fa:"style,omitempty" en:"teash_style,omitempty"`                 // Noe
	StyleEn      string  `gorm:"not null" db:"style" json:"styleEn,omitempty" fa:"teash_style,omitempty" en:"style,omitempty"`                 // Noe
	CodeNaqshe   string  `gorm:"not null" db:"code_naqshe" json:"code_naqshe,omitempty" fa:"code_naghshe,omitempty" en:"code_naghshe,omitempty"`
	CollectionID int64   `gorm:"not null" fa:"-" en:"-"`
	Slug         string  `gorm:"not null" db:"slug" json:"slug,omitempty" fa:"slug,omitempty" en:"slug,omitempty"`
	MostPopular  float64 `gorm:"not null" json:"most_popular,omitempty" fa:"most_popular,omitempty" en:"most_popular,omitempty"`
	BestSelling  float64 `gorm:"not null" json:"best_selling,omitempty" fa:"best_selling,omitempty" en:"best_selling,omitempty"`
	CarpetColors []CarpetColor
}
