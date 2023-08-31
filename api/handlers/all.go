package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hosseinmirzapur/arbitrage/services"
)

func AllUSDT(c *gin.Context) {
	abantether, err := services.RGet("abantether")
	if err != nil {
		handleError(c, err)
		return
	}
	nobitex, err := services.RGet("nobitex")
	if err != nil {
		handleError(c, err)
		return
	}
	okex, err := services.RGet("okex")
	if err != nil {
		handleError(c, err)
		return
	}
	ramzinex, err := services.RGet("ramzinex")
	if err != nil {
		handleError(c, err)
		return
	}
	wallex, err := services.RGet("wallex")
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, &gin.H{
		"abantether": abantether,
		"nobitex":    nobitex,
		"okex":       okex,
		"ramzinex":   ramzinex,
		"wallex":     wallex,
	})
}

func handleError(c *gin.Context, err error) {

	c.JSON(400, &gin.H{
		"message": err.Error(),
	})

}
