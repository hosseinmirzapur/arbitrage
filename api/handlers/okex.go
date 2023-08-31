package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hosseinmirzapur/arbitrage/api/data"
)

func OkexUSDT(c *gin.Context) {

	price, err := data.OkexPrice()

	if err != nil {
		c.JSON(400, &gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(200, &gin.H{
		"okex": &gin.H{
			"buy":  price.Buy,
			"sell": price.Sell,
		},
	})

}
