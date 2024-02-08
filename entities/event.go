package entities

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	CustomerPhone    int64  	`gorm:"column:customer_phone" json:"customer_phone"`
	EventDayWithTime time.Time 	`gorm:"column:event_day_with_time" json:"event_day_with_time"`
	UserStoreId      uint   	`gorm:"column:user_store_id"`
	User			 User		`gorm:"foreignKey:UserStoreId"`
	ServeId        	 uint   	`gorm:"column:serve_id"`
	Service			 Serve		`gorm:"foreignKey:ServeId"`
}