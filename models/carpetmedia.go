package models

type CarpetMedia struct {
	ID            int64  `gorm:"primaryKey;autoIncrement" json:"id" fa:"-" en:"-"`
	Image         string `json:"image" fa:"iamge" en:"iamge"`
	Sort          int64  `json:"sort" fa:"sort,omitempty" en:"sort,omitempty"`
	Feature       string `json:"feature" fa:"feature" en:"feature"` // Background or Original
	CarpetColorID int64  `json:"id_colorcarpet" fa:"-" en:"-"`
}
