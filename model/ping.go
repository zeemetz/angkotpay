package model

import (
	"angkotpay/engine"

	"github.com/jinzhu/gorm"
)

type Ping struct {
	gorm.Model
	Log string
}

func init() {
	engine.GetORM().AutoMigrate(&Ping{})
}
