package services

import (
	"angkotpay/dto/request"
	"angkotpay/dto/response"
	"angkotpay/engine"
	"angkotpay/model"
)

func AuthenticateService(request request.AuthenticateRequest) response.Response {
	var user model.User
	var res response.Response
	user.Phone = request.PhoneNumber
	user.DeviceID = request.DeviceID
	user.Token = request.FirebaseToken
	if engine.GetORM().FirstOrCreate(&user, model.User{Phone: request.PhoneNumber}).Error != nil {
		res.Header = 500
		res.Body = response.ErrorResponse{ErrorCode: 1, Message: "cannot insert or validate request"}
	} else {
		res.Header = 200
		user.Pin = "xxxxxx"
		res.Body = user
	}
	return res
}
