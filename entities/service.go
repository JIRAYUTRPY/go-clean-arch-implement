package entities

import (
	"gorm.io/gorm"
)

type Serve struct {
	gorm.Model
	ServeName     string	`gorm:"column:serve_name" json:"serve_name"`
	ServeDuration int16		`gorm:"column:sserve_duration" json:"serve_duration"`
	UserStoreId     uint	`gorm:"column:user_store_id" json:"user_store_id"`
	ServeColor    string	`gorm:"column:serve_color" json:"serve_color"`
	User			User	`gorm:"foreignKey:UserStoreId"`
}