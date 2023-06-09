package models

type CarpetColor struct {
	ID        int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	ColorCode string `json:"color_code"`
	NameFa    string `json:"color_name_fa"`
	NameEn    string `json:"color_name_en"`
	Default   bool   `json:"default"`
	CarpetID  int64  `gorm:"not null" json:"id_carpet"`
	Carpet    Carpet `gorm:"foreignKey:CarpetID"`
}
