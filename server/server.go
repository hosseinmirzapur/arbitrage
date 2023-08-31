package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hosseinmirzapur/arbitrage/handlers"
)

func Run() {
	r := gin.Default()
	r.GET("/nobitex/usdt", handlers.NobitexUSDT)
	r.GET("/wallex/usdt", handlers.WallexUSDT)
	r.GET("/okex/usdt", handlers.OkexUSDT)
	r.GET("/abantether/usdt", handlers.AbanTetherUSDT)
	r.GET("/ramzinex/usdt", handlers.RamzinexUSDT)
	r.Run(":8000")
}
