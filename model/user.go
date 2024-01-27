package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	//Id      int64  `gorm:"column:id"`
	gorm.Model
	Name    string `json:"name"`
	Email   string `json:"email"`
	Country string `json:"country"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type PostUser struct {
	Name    string `json:"name,omitempty"`
	Email   string `json:"email,omitempty"`
	Country string `json:"country,omitempty"`
	Phone   string `json:"phone,omitempty"`
	Address string `json:"address,omitempty"`
}

type Article struct {
	gorm.Model
	Title     string    `json:"title"`
	Subtitle  string    `json:"subtitle"`
	Author    string    `json:"author"`
	Upcounter uint      `json:"upcounter"`
	Postdate  time.Time `json:"postdate"`
}
