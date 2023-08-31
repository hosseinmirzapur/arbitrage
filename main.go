package main

import (
	"log"

	"github.com/hosseinmirzapur/arbitrage/server"
	"github.com/hosseinmirzapur/arbitrage/services"
	"github.com/hosseinmirzapur/arbitrage/services/jobs"
)

func main() {
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
