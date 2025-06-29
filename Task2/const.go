package main

import (
	"fmt"
	"math"
)

func main() {

	const (
		Pi = math.Pi
		e  = math.E
	)

	fmt.Println("Число PI = ", Pi)
	fmt.Println("Число E = ", e)
	fmt.Println("Длина окруж-ти при радиусе (R=4)=", 2*Pi*4)
	fmt.Println("Логариф по основанию E =", math.Log(e))
}
