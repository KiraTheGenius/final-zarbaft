package models

type CarpetColor struct {
	ID           int64         `gorm:"primaryKey;autoIncrement" json:"id " fa:"-" en:"-"`
	ColorCode    string        `json:"color_code" fa:"color_code,omitempty" en:"color_code,omitempty"`
	NameFa       string        `json:"color_name_fa" fa:"name,omitempty" en:"trash_name,omitempty"`
	NameEn       string        `json:"color_name_en" fa:"trash_name,omitempty" en:"name,omitempty"`
	Default      bool          `json:"default" fa:"defult" en:"defult"`
	CarpetID     int64         `gorm:"not null" json:"carpet_id" fa:"-" en:"-"`
	CarpetMedias []CarpetMedia `fa:"media" en:"media"`
}
