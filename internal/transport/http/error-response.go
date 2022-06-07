package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

type errorResponse struct {
	ErrorTime string `json:"errorTime"`
	Message   string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{
		ErrorTime: time.Now().Format(timeFormat),
		Message:   message,
	})
}
