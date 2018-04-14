package model

import (
	"angkotpay/engine"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Phone      string
	DeviceID   string
	Token      string
	Name       string
	Pin        string
	Balance    int64 `gorm:"default:0"`
	UpperLimit int64 `gorm:"default:1000000"`
	LowerLimit int64 `gorm:"default:0"`
}

func init() {
	engine.GetORM().AutoMigrate(&User{})
}
