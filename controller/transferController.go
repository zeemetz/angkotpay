package controller

import (
	"angkotpay/dto/request"
	"angkotpay/services"

	"github.com/gin-gonic/gin"
)

func Transfer(c *gin.Context) {
	var request request.TransferRequest
	if c.Bind(&request) != nil {
		c.JSON(400, "bad request")
	} else {
		res := services.TransferService(request)
		c.JSON(res.Header, res.Body)
	}
}
