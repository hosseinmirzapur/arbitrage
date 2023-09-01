package main

import (
	"log"

	"github.com/hosseinmirzapur/arbitrage/config"
	"github.com/hosseinmirzapur/arbitrage/server"
	"github.com/hosseinmirzapur/arbitrage/services"
	"github.com/hosseinmirzapur/arbitrage/services/jobs"
	"github.com/joho/godotenv"
)

func main() {

	serverConf := config.NewConf()
	if serverConf.AppMode() == "prod" {
		godotenv.Load()
	}

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

	// Start Server
	err = server.Run()
	if err != nil {
		log.Fatalf("cannot start the server...\n")
	}
}
