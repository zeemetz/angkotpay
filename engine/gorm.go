package engine

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var database *gorm.DB

func openDatabase() {
	var err error
	pwd, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(pwd)

	// database, err = gorm.Open("sqlite3", pwd+"/angkotpay.db")
	database, err = gorm.Open("postgres", "host=128.199.184.58 port=5432 user=postgres dbname=angkotpay password=evo123 sslmode=disable")

	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	database.DB().SetMaxOpenConns(100)
	database.DB().SetMaxIdleConns(1)
}

func init() {
	openDatabase()
}

//GetORM is
func GetORM() *gorm.DB {
	if database == nil {
		openDatabase()
	}
	return database.Debug()
}
