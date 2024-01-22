package v1

import (
	"github.com/gin-gonic/gin"
	"userLoggingProject/pkg/logger"
)

type dataResponse struct {
	Data interface{} `json:"data"`
}

type response struct {
	Message string `json:"message"`
}

func newResponse(c *gin.Context, statusCode int, message string) {
	logger.Errorf(message)
	c.AbortWithStatusJSON(statusCode, response{message})
}
