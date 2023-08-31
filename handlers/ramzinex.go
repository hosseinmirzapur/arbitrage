package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hosseinmirzapur/arbitrage/config"
	"github.com/hosseinmirzapur/arbitrage/utils"
)

func RamzinexUSDT(c *gin.Context) {
	endpoint := config.Ramzinex().MarketURL

	data, err := utils.GetRequest(endpoint)

	if err != nil {
		c.JSON(400, &gin.H{
			"message": err.Error(),
		})
		return
	}

	mainData := data["data"].(map[string]interface{})

	buy := mainData["buy"].(float64)
	sell := mainData["sell"].(float64)

	c.JSON(200, &gin.H{
		"ramzinex": &gin.H{
			"buy":  buy / 10,
			"sell": sell / 10,
		},
	})
}
