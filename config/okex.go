package config

type okexConfig struct {
	exchangeConfig
}

func Okex() *okexConfig {
	return &okexConfig{
		exchangeConfig: exchangeConfig{
			MarketURL: "https://api.ok-ex.io/oapi/v1/market/tickers",
			RialAbbr:  "IRT",
		},
	}
}
