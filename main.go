package main

import (
	"github.com/hosseinmirzapur/arbitrage/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	server.Run()
}
