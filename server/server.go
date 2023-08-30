package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hosseinmirzapur/arbitrage/services"
)

func Run() {
	// r := gin.New()
	// r.Use(gin.Recovery())
	r := gin.Default()
	r.GET("/nobitex/usdt", services.NobitexUSDT)
	r.GET("/wallex/usdt", services.WallexUSDT)
	r.GET("/okex/usdt", services.OkexUSDT)
	r.Run(":3000")
}
