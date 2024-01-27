package model

import "gorm.io/gorm"

type User struct {
	//Id      int64  `gorm:"column:id"`
	gorm.Model
	Name    string `gorm:"column:name"`
	Email   string `gorm:"column:email"`
	Country string `gorm:"column:country"`
	Phone   string `gorm:"column:phone"`
	Address string `gorm:"column:address"`
}
