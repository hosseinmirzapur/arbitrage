package data

import (
	"fmt"

	"github.com/hosseinmirzapur/arbitrage/config"
	"github.com/hosseinmirzapur/arbitrage/services/client"
	"github.com/hosseinmirzapur/arbitrage/utils"
)

func NobitexPrice() (*Price, error) {
	usdtPriceURL := fmt.Sprintf("%s/%s%s", config.Nobitex().MarketURL, "USDT", config.Nobitex().RialAbbr)
	data, err := client.GetRequest(usdtPriceURL)

	if err != nil {
		return nil, err
	}

	bids := data["bids"].([]any)
	asks := data["asks"].([]any)

	buyPrice := bids[0].([]any)[0].(string)
	sellPrice := asks[0].([]any)[0].(string)

	return &Price{
		Buy:  utils.StringToFloat(buyPrice) / 10,
		Sell: utils.StringToFloat(sellPrice) / 10,
	}, nil

}
