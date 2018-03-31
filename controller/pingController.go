package controller

import (
	"angkotpay/services"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	str := c.Query("str")
	res := services.Ping(str)
	c.JSON(res.Header, res.Body)
}
