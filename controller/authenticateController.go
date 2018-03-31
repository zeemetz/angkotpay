package controller

import (
	"angkotpay/dto/request"
	"angkotpay/services"

	"github.com/gin-gonic/gin"
)

func AuthenticateController(c *gin.Context) {
	var request request.AuthenticateRequest
	c.Bind(&request)
	res := services.AuthenticateService(request)
	c.JSON(res.Header, res.Body)
}
