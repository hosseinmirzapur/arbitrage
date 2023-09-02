package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hosseinmirzapur/arbitrage/api/handlers"
	"github.com/hosseinmirzapur/arbitrage/config"
)

func Run(conf *config.BaseConfig) error {
	r := gin.Default()

	// r.SetTrustedProxies([]string{"http://localhost:3000"})
	r.GET("/nobitex/usdt", handlers.NobitexUSDT)
	r.GET("/wallex/usdt", handlers.WallexUSDT)
	r.GET("/okex/usdt", handlers.OkexUSDT)
	r.GET("/abantether/usdt", handlers.AbanTetherUSDT)
	r.GET("/ramzinex/usdt", handlers.RamzinexUSDT)
	r.GET("/all/usdt", handlers.AllUSDT)
	return run(r, conf)
}

func run(r *gin.Engine, conf *config.BaseConfig) error {
	if conf.AppMode() == "prod" {
		return r.RunTLS(":3000", "server.crt", "server.key")
	}
	return r.Run(":3000")
}
