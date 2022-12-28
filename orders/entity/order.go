package entity

// 订单实体
type Order struct {
	Id        uint    `json:"id"  gorm:"id"`
	GoodId    uint    `json:"goodId"  gorm:"goodId"`
	GoodCount uint    `json:"goodCount"  gorm:"goodCount"`
	Sum       float64 `json:"sum"  gorm:"sum"`
}
