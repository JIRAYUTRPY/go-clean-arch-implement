package entities

import (
	"gorm.io/gorm"
)

type Serve struct {
	gorm.Model
	Name     		string	`gorm:"column:name" json:"serve_name"`
	Duration 		int16	`gorm:"column:duration" json:"serve_duration"`
	Color    		string	`gorm:"column:color" json:"serve_color"`
	UserId			uint	`gormm:"column:user_id" json:"user_id"`
}