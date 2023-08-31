package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hosseinmirzapur/arbitrage/config"
	"github.com/hosseinmirzapur/arbitrage/utils"
)

func WallexUSDT(c *gin.Context) {
	dataURL := config.Wallex().MarketURL

	data, err := utils.GetRequest(dataURL)

	if err != nil {
		log.Println(err.Error())
	}

	result := data["result"].(map[string]interface{})

	symbols := result["symbols"].(map[string]interface{})

	for index, item := range symbols {
		if index == "USDTTMN" {
			stats := item.(map[string]interface{})["stats"].(map[string]interface{})

			sellPrice := stats["askPrice"].(string)
			buyPrice := stats["bidPrice"].(string)

			c.JSON(200, &gin.H{
				"wallex": &gin.H{
					"buy":  utils.StringToFloat(buyPrice),
					"sell": utils.StringToFloat(sellPrice),
				},
			})
			return
		}
	}

	c.JSON(404, &gin.H{
		"message": "Not Found",
	})
}
