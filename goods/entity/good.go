package entity

type Good struct {
	Id       uint    `json:"id"  gorm:"id"`
	GoodName string  `json:"goodname"  gorm:"goodname"`
	Price    float64 `json:"price"  gorm:"price"`
}
