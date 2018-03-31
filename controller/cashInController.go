package controller

import (
	"angkotpay/dto/request"
	"angkotpay/services"

	"github.com/gin-gonic/gin"
)

func CashIn(c *gin.Context) {
	var request request.TransferRequest
	c.Bind(&request)
	res := services.TransferService(request)
	c.JSON(res.Header, res.Body)
}
