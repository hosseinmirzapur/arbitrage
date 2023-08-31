package data

import (
	"github.com/hosseinmirzapur/arbitrage/config"
	"github.com/hosseinmirzapur/arbitrage/utils"
)

func AbanPrice() (*Price, error) {
	endpoint := config.AbanTether().MarketURL + "?coin=USDT"

	data, err := utils.AbanGetRequest(endpoint)
	if err != nil {
		return nil, err
	}

	usdt := data["USDT"].(map[string]interface{})
	buyPrice := usdt["irtPriceBuy"]
	sellPrice := usdt["irtPriceSell"]

	return &Price{
		Buy:  utils.StringToFloat(buyPrice.(string)),
		Sell: utils.StringToFloat(sellPrice.(string)),
	}, nil
}
