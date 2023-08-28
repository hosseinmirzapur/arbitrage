package services

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hosseinmirzapur/arbitrage/config"
	"github.com/hosseinmirzapur/arbitrage/utils"
)

func NobitexUSDT(c *gin.Context) {
	usdtPriceURL := fmt.Sprintf("%s/%s%s", config.Nobitex().MarketURL, "USDT", config.Nobitex().RialAbbr)
	data, err := utils.GetRequest(usdtPriceURL)

	if err != nil {
		log.Println(err)
	}

	if err != nil {
		log.Println(err)
	}

	asks := data["asks"].([]any)
	bids := data["bids"].([]any)

	buyPrice := bids[0].([]any)[0].(string)
	sellPrice := asks[0].([]any)[0].(string)

	c.JSON(200, &gin.H{
		"nobitex": &gin.H{
			"buy":  utils.StringToInt(buyPrice) / 10,
			"sell": utils.StringToInt(sellPrice) / 10,
		},
	})
}
