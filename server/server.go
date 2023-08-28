package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hosseinmirzapur/arbitrage/services"
)

func Run() {
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/nobitex/usdt", services.NobitexUSDT)
	r.GET("/wallex/usdt", services.WallexUSDT)
	r.Run(":3000")
}
