package main

import "fmt"

type Car struct {
	Type    string
	FuelInL float64
}

// Method untuk menghitung perkiraan jarak yang bisa ditempuh berdasarkan bahan bakar
func (c Car) CalculateEstimatedDistance() float64 {
	// Asumsi efisiensi bahan bakar 1.5 L/mil
	efficiency := 1.5
	return c.FuelInL / efficiency
}

func main() {
	// Membuat instance Car
	myCar := Car{
		Type:    "SUV",
		FuelInL: 45.0,
	}

	// Menghitung perkiraan jarak yang bisa ditempuh
	estimatedDistance := myCar.CalculateEstimatedDistance()
	fmt.Printf("Perkiraan jarak yang bisa ditempuh oleh mobil %s adalah %.2f mil.\n", myCar.Type, estimatedDistance)
}
