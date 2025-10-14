package utils

import "github.com/gin-gonic/gin"

func FormatErrorBody(ctx *gin.Context, statusCode int, errorMessage string) {
	ctx.JSON(statusCode, gin.H{
		"error": errorMessage,
	})
}
