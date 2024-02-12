package entities

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	CustomerPhone    string  	`gorm:"column:phone" json:"customer_phone"`
	CustomerName	 string		`gorm:"column:name" json:"customer_name"`
	EventDayWithTime string  	`gorm:"column:event_day_with_time" json:"event_day_with_time"`
	UserId		 	 uint		`gorm:"column:user_id"`
	ServeId			 uint		`gorm:"column:serve_id"`
}