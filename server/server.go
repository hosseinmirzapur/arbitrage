package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hosseinmirzapur/arbitrage/api/handlers"
)

func Run() error {
	r := gin.Default()
	r.GET("/nobitex/usdt", handlers.NobitexUSDT)
	r.GET("/wallex/usdt", handlers.WallexUSDT)
	r.GET("/okex/usdt", handlers.OkexUSDT)
	r.GET("/abantether/usdt", handlers.AbanTetherUSDT)
	r.GET("/ramzinex/usdt", handlers.RamzinexUSDT)
	r.GET("/all/usdt", handlers.AllUSDT)
	return r.Run(":8000")
}
