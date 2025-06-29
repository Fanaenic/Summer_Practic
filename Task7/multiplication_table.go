package main

import "fmt"

func main() {

	fmt.Println("Произведение чисел от 1 до 10\n")

	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
			fmt.Print("\n")
		}
	}
}
