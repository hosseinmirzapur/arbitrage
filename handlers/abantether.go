package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hosseinmirzapur/arbitrage/config"
	"github.com/hosseinmirzapur/arbitrage/utils"
)

func AbanTetherUSDT(c *gin.Context) {
	endpoint := config.AbanTether().MarketURL + "?coin=USDT"

	data, err := utils.AbanGetRequest(endpoint)
	if err != nil {
		log.Println(err)
	}

	usdt := data["USDT"].(map[string]interface{})
	buyPrice := usdt["irtPriceBuy"]
	sellPrice := usdt["irtPriceSell"]

	c.JSON(200, &gin.H{
		"abantether": &gin.H{
			"buy":  buyPrice,
			"sell": sellPrice,
		},
	})

}
