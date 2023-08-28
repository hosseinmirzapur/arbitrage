package utils

func CalculateProfit(buyPrice, sellPrice float64, quantity int) float64 {
	return (sellPrice - buyPrice) * float64(quantity)
}
