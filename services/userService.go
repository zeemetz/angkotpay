package services

import (
	"angkotpay/dto/response"
	"angkotpay/engine"
	"angkotpay/model"
)

func UserService(user model.User) response.Response {
	var res response.Response
	if engine.GetORM().Find(&user, model.User{Phone: user.Phone}).Error != nil {
		res.Header = 404
		res.Body = response.ErrorResponse{ErrorCode: 8, Message: "user not found"}
	} else if engine.GetORM().Model(&user).Updates(user).Error != nil {
		res.Header = 400
		res.Body = response.ErrorResponse{ErrorCode: 8, Message: "invalid user updates"}
	} else {
		res.Header = 200
		user.Pin = "xxxxxx"
		res.Body = user
	}

	return res
}
