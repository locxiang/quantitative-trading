package controllers

import (
	"github.com/gin-gonic/gin"
	"gitlab.singulato.com/neltharion/virtual-car/app/errno"
)


type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}, status int) {
	code, message := errno.DecodeErr(err)
	// always return http.StatusOK
	c.JSON(status, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
