package jobs

import (
	"time"

	"github.com/go-co-op/gocron"
)

func RunCron() error {
	s := gocron.NewScheduler(time.UTC)
	s.SingletonModeAll()

	_, err := s.Every(5).Minutes().Do(updatePrices)

	return err

}
