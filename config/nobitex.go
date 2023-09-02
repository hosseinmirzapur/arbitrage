package config

type nobitexConfig struct {
	exchangeConfig
}

func Nobitex() *nobitexConfig {
	return &nobitexConfig{
		exchangeConfig: exchangeConfig{
			MarketURL: "https://api.nobitex.ir/market/stats",
			RialAbbr:  "rls",
		},
	}
}
