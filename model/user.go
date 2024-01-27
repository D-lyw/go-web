package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	//Id      int64  `gorm:"column:id"`
	gorm.Model
	Name    string
	Email   string
	Country string
	Phone   string
	Address string
}

type Article struct {
	gorm.Model
	Title     string
	Subtitle  string
	Author    string
	Upcounter uint
	Postdate  time.Time
}
