package config

type wallexConfig struct {
	ExchangeConfig
}

func Wallex() *wallexConfig {
	return &wallexConfig{
		ExchangeConfig: ExchangeConfig{
			MarketURL: "https://api.wallex.ir/v1/markets",
			RialAbbr:  "TMN",
		},
	}
}
