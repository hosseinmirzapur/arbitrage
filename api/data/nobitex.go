package data

import (
	"github.com/hosseinmirzapur/arbitrage/config"
	"github.com/hosseinmirzapur/arbitrage/services/client"
	"github.com/hosseinmirzapur/arbitrage/utils"
)

func NobitexPrice() (*Price, error) {
	usdtPriceURL := config.Nobitex().MarketURL
	body := []byte(`{
		"srcCurrency": "usdt",
		"dstCurrency": "rls"
	}`)

	data, err := client.PostRequest(usdtPriceURL, body)

	if err != nil {
		return nil, err
	}

	stats := data["stats"].(map[string]interface{})
	usdtRls := stats["usdt-rls"].(map[string]interface{})

	buyPrice := usdtRls["bestBuy"].(string)
	sellPrice := usdtRls["bestSell"].(string)

	return &Price{
		Buy:  utils.StringToFloat(buyPrice) / 10,
		Sell: utils.StringToFloat(sellPrice) / 10,
	}, nil

}
