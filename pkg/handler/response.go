package handler

import "github.com/gin-gonic/gin"

// responseWithError отдает ответ с ошибкой и JSON
func responseWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

// responseSuccessful отдает 200 ответ с JSON
func responseSuccessful(c *gin.Context, message interface{}) {
	c.AbortWithStatusJSON(200, gin.H{"successful": message})
}
