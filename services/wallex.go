package services

import "github.com/gin-gonic/gin"

func WallexUSDT(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Working fine",
	})
}
