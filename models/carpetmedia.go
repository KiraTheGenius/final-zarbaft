package models

type CarpetMedia struct {
	ID            int64       `gorm:"primaryKey;autoIncrement" json:"id"`
	Image         string      `json:"image"`
	Sort          int64       `json:"sort"`
	Feature       string      `json:"feature"` // Background or Original
	CarpetColorID int64       `json:"id_colorcarpet"`
	CarpetColor   CarpetColor `gorm:"not null" gorm:"foreignKey:CarpetColorID"`
	CarpetID      int64       `json:"id_carpet"`
	Carpet        Carpet      ` gorm:"foreignKey:CarpetID"`
}
