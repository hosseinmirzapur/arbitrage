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
	price, err := data.AbanPrice()
	handleError(err)

	err = services.RSet("nobitex", price)
	handleError(err)
}

func updateOkex() {
	price, err := data.OkexPrice()
	handleError(err)

	err = services.RSet("okex", price)
	handleError(err)
}

func updateRamzinex() {
	price, err := data.RamzinexPrice()
	handleError(err)

	err = services.RSet("ramzinex", price)
	handleError(err)
}

func updateWallex() {
	price, err := data.WallexPrice()
	handleError(err)

	err = services.RSet("wallex", price)
	handleError(err)
}

func updateAbanTether() {
	price, err := data.AbanPrice()
	handleError(err)

	err = services.RSet("abantether", price)
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		log.Println(err)
	}
}
