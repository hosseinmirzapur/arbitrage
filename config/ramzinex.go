package config

type ramzinexConfig struct {
	exchangeConfig
}

func Ramzinex() *ramzinexConfig {
	return &ramzinexConfig{
		exchangeConfig: exchangeConfig{
			MarketURL: "https://publicapi.ramzinex.com/exchange/api/v1.0/exchange/pairs/11",
			RialAbbr:  "irr",
		},
	}
}
