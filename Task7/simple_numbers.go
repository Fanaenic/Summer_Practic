package main

import (
	"fmt"
	"math"
)

func main() {
	for num := 2; num <= 100; num++ {

		isPrime := true

		sqrt := int(math.Sqrt(float64(num))) + 1

		for i := 2; i < sqrt; i++ {

			if num%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(num)
		}
	}
}
