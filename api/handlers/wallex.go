package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hosseinmirzapur/arbitrage/api/data"
)

func WallexUSDT(c *gin.Context) {

	price, err := data.WallexPrice()

	if err != nil {
		c.JSON(400, &gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(200, &gin.H{
		"wallex": &gin.H{
			"buy":  price.Buy,
			"sell": price.Sell,
		},
	})

}
