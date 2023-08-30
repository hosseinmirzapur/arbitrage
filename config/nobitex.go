package config

type nobitexConfig struct {
	exchangeConfig
}

func Nobitex() *nobitexConfig {
	return &nobitexConfig{
		exchangeConfig: exchangeConfig{
			MarketURL: "https://api.nobitex.ir/v2/orderbook",
			RialAbbr:  "IRT",
		},
	}
}
