package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hosseinmirzapur/arbitrage/api/data"
)

func RamzinexUSDT(c *gin.Context) {

	price, err := data.RamzinexPrice()
	if err != nil {
		c.JSON(400, &gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, &gin.H{
		"ramzinex": &gin.H{
			"buy":  price.Buy,
			"sell": price.Sell,
		},
	})
}
