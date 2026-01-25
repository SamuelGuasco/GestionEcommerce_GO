package main

func CalcularTotal(precios []float64) float64 {
	total := 0.0
	for _, p := range precios {
		total += p
	}
	return total
}
