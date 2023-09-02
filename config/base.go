package config

type baseConfig struct {
	appMode string
}

func NewConf() *baseConfig {
	return &baseConfig{
		appMode: "prod", // "dev" or "prod"
	}
}

func (c *baseConfig) AppMode() string {
	return c.appMode
}
