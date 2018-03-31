package controller

import (
	"angkotpay/model"
	"angkotpay/services"

	"github.com/gin-gonic/gin"
)

func UserController(c *gin.Context) {
	var user model.User
	c.Bind(&user)
	res := services.UserService(user)
	c.JSON(res.Header, res.Body)
}
