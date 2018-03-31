package main

import (
	"angkotpay/controller"
	"angkotpay/engine"
	"os"
)

func init() {
	engine.GetAdmin().GET("/ping", controller.Ping)
	engine.GetEngine().POST("/authenticate", controller.AuthenticateController)
	// engine.GetEngine().POST("/cashin", controller.CashIn)
	engine.GetEngine().POST("/transfer", controller.Transfer)
}

func main() {
	if len(os.Args) < 3 {
		engine.GetEngine().Run("localhost:1313")
	}
	engine.GetEngine().Run(os.Args[1] + ":" + os.Args[2])
}
