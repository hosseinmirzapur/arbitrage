package services

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hosseinmirzapur/arbitrage/config"
	"github.com/hosseinmirzapur/arbitrage/utils"
)

func OkexUSDT(c *gin.Context) {
	endpoint := config.Okex().MarketURL

	data, err := utils.GetRequest(endpoint)

	if err != nil {
		log.Println(err)
	}

	ticker := data["tickers"].([]interface{})

	for _, v := range ticker {
		if item := v.(map[string]interface{}); item["symbol"] == "USDT-IRT" {
			buyPrice := item["best_bid"].(string)
			sellPrice := item["best_ask"].(string)

			c.JSON(200, &gin.H{
				"okex": &gin.H{
					"buy":  utils.StringToFloat(buyPrice),
					"sell": utils.StringToFloat(sellPrice),
				},
			})
			return
		}
	}
	c.JSON(404, &gin.H{
		"message": "market not found",
	})

}
