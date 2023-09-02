package data

import (
	"github.com/hosseinmirzapur/arbitrage/config"
	"github.com/hosseinmirzapur/arbitrage/services/client"
)

func RamzinexPrice() (*Price, error) {
	endpoint := config.Ramzinex().MarketURL

	data, err := client.GetRequest(endpoint)

	if err != nil {
		return nil, err
	}

	mainData := data["data"].(map[string]interface{})

	buy := mainData["buy"].(float64)
	sell := mainData["sell"].(float64)

	return &Price{
		Buy:  buy / 10,
		Sell: sell / 10,
	}, nil
}
