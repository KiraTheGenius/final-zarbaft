package models

type Collection struct {
	ID         int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	NameFa     string `gorm:"not null" json:"name_fa"`
	MottoFa    string `gorm:"not null" json:"motto_fa"` // Shoaar
	NameEn     string `gorm:"not null" json:"name_en"`
	MottoEn    string `gorm:"not null" json:"motto_en"` // Shoaar
	Slug       string `gorm:"not null" json:"slug"`
	Background string `gorm:"not null" json:"pic1"` // Images for Background (1)
	Collection string `gorm:"not null" json:"pic2"` // Images for Collection (2)
	Carpets    []Carpet
	// Images for Background and Collection (2)
}
