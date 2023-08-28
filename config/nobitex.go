package config

type NobitexConfig struct {
	ExchangeConfig
}

func Nobitex() *NobitexConfig {
	return &NobitexConfig{
		ExchangeConfig: ExchangeConfig{
			MarketURL: "https://api.nobitex.ir/v2/orderbook",
			RialAbbr:  "IRT",
		},
	}
}
