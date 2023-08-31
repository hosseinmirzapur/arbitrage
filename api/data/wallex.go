package data

import (
	"errors"

	"github.com/hosseinmirzapur/arbitrage/config"
	"github.com/hosseinmirzapur/arbitrage/utils"
)

func WallexPrice() (*Price, error) {
	dataURL := config.Wallex().MarketURL

	data, err := utils.GetRequest(dataURL)

	if err != nil {
		return nil, err
	}

	result := data["result"].(map[string]interface{})

	symbols := result["symbols"].(map[string]interface{})

	for index, item := range symbols {
		if index == "USDTTMN" {
			stats := item.(map[string]interface{})["stats"].(map[string]interface{})

			sellPrice := stats["askPrice"].(string)
			buyPrice := stats["bidPrice"].(string)

			return &Price{
				Buy:  utils.StringToFloat(buyPrice),
				Sell: utils.StringToFloat(sellPrice),
			}, nil
		}
	}

	return nil, errors.New("price not found")
}
