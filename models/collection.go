package models

type Collection struct {
	ID         int64    `gorm:"primaryKey;autoIncrement" json:"id" fa:"-" en:"-"`
	NameFa     string   `gorm:"not null" json:"name_fa" fa:"name,omitempty" en:"trash_name,omitempty"`
	MottoFa    string   `gorm:"not null" json:"motto_fa" fa:"motto,omitempty" en:"trash_motto,omitempty"` // Shoaar
	NameEn     string   `gorm:"not null" json:"name_en" fa:"trash_name,omitempty" en:"name,omitempty"`
	MottoEn    string   `gorm:"not null" json:"motto_en" fa:"trash_motto,omitempty" en:"motto,omitempty"` // Shoaar
	Slug       string   `gorm:"not null" json:"slug" fa:"slug,omitempty" en:"slug,omitempty"`
	Background string   `gorm:"not null" json:"pic1" fa:"pic1,omitempty" en:"pic1,omitempty"` // Images for Background (1)
	Collection string   `gorm:"not null" json:"pic2" fa:"pic2,omitempty" en:"pic2,omitempty"` // Images for Collection (2)
	Carpets    []Carpet `fa:"carpets"`
	// Images for Background and Collection (2)
}
