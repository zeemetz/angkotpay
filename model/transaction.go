package model

import (
	"angkotpay/engine"

	"github.com/jinzhu/gorm"
)

const (
	TRANSFER = "TRANSFER"
	CASHIN   = "CASHIN"
	FAILED   = "FAILED"
	SUCCESS  = "SUCCESS"
)

type Transaction struct {
	gorm.Model
	From            uint
	To              uint
	Amount          int64
	Notes           string
	TransactionType string
	Status          string
	LastBalance     int64
}

func init() {
	engine.GetORM().AutoMigrate(&Transaction{})
}
