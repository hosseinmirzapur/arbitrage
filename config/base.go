package config

type baseConfig struct {
	appMode string
}

func NewConf() *baseConfig {
	return &baseConfig{
		appMode: "dev", // "dev" or "prod"
	}
}

func (c *baseConfig) AppMode() string {
	return c.appMode
}
