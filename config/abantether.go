package config

import "os"

type abanTetherConfig struct {
	exchangeConfig
	apiKey string
}

func AbanTether() *abanTetherConfig {

	apiKey := os.Getenv("ABANTETHER_API_KEY")

	return &abanTetherConfig{
		exchangeConfig: exchangeConfig{
			MarketURL: "https://abantether.com/api/v1/otc/coin-price",
			RialAbbr:  "irt",
		},
		apiKey: apiKey,
	}
}
