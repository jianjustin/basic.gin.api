package model

type SysUser struct {
	Id       uint   `json:"id"  gorm:"id"`
	UserName string `json:"username"  gorm:"username"`
	Password string `json:"password"  gorm:"password"`
}
