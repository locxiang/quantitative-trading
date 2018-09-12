package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/locxiang/quantitative-trading/app/errors"
)


type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err *errors.HttpError, data interface{}, status int) {
	// always return http.StatusOK
	c.JSON(status, Response{
		Code:    err.Code,
		Message: err.Message,
		Data:    data,
	})
}
