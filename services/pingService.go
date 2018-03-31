package services

import (
	"angkotpay/dto/response"
	"angkotpay/engine"
	"angkotpay/model"
)

func Ping(str string) response.Response {
	var ping model.Ping
	ping.Log = str
	engine.GetORM().Create(&ping)
	return response.Response{Header: 200, Body: "hello world " + str}
}
