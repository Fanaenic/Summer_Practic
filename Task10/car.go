package main

import "fmt"

type Engine struct {
	Power    int
	Volume   float32
	FuelType string
}

type Car struct {
	Brand string
	Model string
	Year  int
	Engine
}

func main() {

	zhiga := Car{
		Brand: "ВАЗ",
		Model: "Жигули 2106",
		Year:  1985,
		Engine: Engine{
			Power:    75,
			Volume:   1.6,
			FuelType: "АИ-92",
		},
	}

	printCarInfo(zhiga)
}

func printCarInfo(c Car) {
	fmt.Println("Легендарный автомобиль:")
	fmt.Printf("%s %s %d года выпуска\n", c.Brand, c.Model, c.Year)
	fmt.Printf("Двигатель: %.1f л, %d л.с., бензин %s\n",
		c.Volume, c.Power, c.FuelType)
}
