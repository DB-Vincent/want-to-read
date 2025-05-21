package models

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	IsSuper  bool   `json:"is_super" gorm:"default:false"`
}
