package libs

import "github.com/gin-gonic/gin"

func ErrResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"code":    statusCode,
		"message": message,
	})
}
