package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hosseinmirzapur/arbitrage/config"
	"github.com/hosseinmirzapur/arbitrage/server"
	"github.com/hosseinmirzapur/arbitrage/services"
	"github.com/hosseinmirzapur/arbitrage/services/jobs"
	"github.com/joho/godotenv"
)

func main() {

	serverConf := config.NewConf()
	if serverConf.AppMode() == "dev" {
		godotenv.Load()
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
		serverMethods()
	}

	// Start Server
	err := server.Run(serverConf)
	if err != nil {
		log.Fatalf("cannot start the server...\n")
	}
}

func serverMethods() {
	// Redis Init
	err := services.RedisInit()
	if err != nil {
		log.Fatalf("cannot start redis client...\n")
	}

	// Cron Jobs
	err = jobs.RunCron()
	if err != nil {
		log.Fatalf("cannot run cron jobs...\n")
	}
}
