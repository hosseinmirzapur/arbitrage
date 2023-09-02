package jobs

import (
	"log"

	"github.com/hosseinmirzapur/arbitrage/api/data"
	"github.com/hosseinmirzapur/arbitrage/services"
)

func updatePrices() {
	go updateNobitex()
	go updateOkex()
	go updateRamzinex()
	go updateWallex()
	go updateAbanTether()
}

func updateNobitex() {
	log.Println("updating nobitex...")
	price, err := data.AbanPrice()
	handleError(err)

	err = services.RSet("nobitex", price)
	handleError(err)
	log.Println("updated nobitex")
}

func updateOkex() {
	log.Println("updating okex...")
	price, err := data.OkexPrice()
	handleError(err)

	err = services.RSet("okex", price)
	handleError(err)
	log.Println("updated okex")
}

func updateRamzinex() {
	log.Println("updating ramzinex...")
	price, err := data.RamzinexPrice()
	handleError(err)

	err = services.RSet("ramzinex", price)
	handleError(err)
	log.Println("updated ramzinex")
}

func updateWallex() {
	log.Println("updating wallex...")
	price, err := data.WallexPrice()
	handleError(err)

	err = services.RSet("wallex", price)
	handleError(err)
	log.Println("updated wallex")
}

func updateAbanTether() {
	log.Println("updating abantether...")
	price, err := data.AbanPrice()
	handleError(err)

	err = services.RSet("abantether", price)
	handleError(err)
	log.Println("updated abantether")
}

func handleError(err error) {
	if err != nil {
		log.Println(err)
	}
}
