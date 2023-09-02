package data

import (
	"errors"

	"github.com/hosseinmirzapur/arbitrage/config"
	"github.com/hosseinmirzapur/arbitrage/services/client"
	"github.com/hosseinmirzapur/arbitrage/utils"
)

func OkexPrice() (*Price, error) {
	endpoint := config.Okex().MarketURL

	data, err := client.GetRequest(endpoint)

	if err != nil {
		return nil, err
	}

	ticker := data["tickers"].([]interface{})

	for _, v := range ticker {
		if item := v.(map[string]interface{}); item["symbol"] == "USDT-IRT" {
			buyPrice := item["best_bid"].(string)
			sellPrice := item["best_ask"].(string)

			return &Price{
				Buy:  utils.StringToFloat(buyPrice),
				Sell: utils.StringToFloat(sellPrice),
			}, nil
		}
	}

	return nil, errors.New("price not found")
}
