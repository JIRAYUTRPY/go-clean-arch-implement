package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string  `gorm:"column:email" json:"email"`
	Password     string	 `gorm:"column:password" json:"password"`
	Phone        string  `gorm:"column:phone" json:"phone"`
	Address      string  `gorm:"column:address" json:"address"`
	LocationLink string  `gorm:"column:location_link" json:"location_link"`
	StoreName    string  `gorm:"column:store_name" json:"store_name"`
	Role 		 string  `gorm:"column:role" json:"role"`
	Events       []Event
	Serves	     []Serve
}