package config

type BaseConfig struct {
	appMode string
}

func NewConf() *BaseConfig {
	return &BaseConfig{
		appMode: "dev", // "dev" or "prod"
	}
}

func (c *BaseConfig) AppMode() string {
	return c.appMode
}
