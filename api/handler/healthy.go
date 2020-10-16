package handler

import "github.com/gin-gonic/gin"

// HealthyRouters check the healthy of api
func HealthyRouters(e *gin.Engine) {
	e.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "go transactions api is healthy",
		})
	})
}
