package config

type wallexConfig struct {
	exchangeConfig
}

func Wallex() *wallexConfig {
	return &wallexConfig{
		exchangeConfig: exchangeConfig{
			MarketURL: "https://api.wallex.ir/v1/markets",
			RialAbbr:  "TMN",
		},
	}
}
