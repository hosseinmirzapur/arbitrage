package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hosseinmirzapur/arbitrage/api/data"
)

func AbanTetherUSDT(c *gin.Context) {

	price, err := data.AbanPrice()

	if err != nil {
		c.JSON(400, &gin.H{
			"message": err.Error(),
		})
		return
	}
	buyPrice := price.Buy
	sellPrice := price.Sell

	c.JSON(200, &gin.H{
		"abantether": &gin.H{
			"buy":  buyPrice,
			"sell": sellPrice,
		},
	})

}
