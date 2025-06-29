package main

import "fmt"

func incrementByValue(a int) {
	a = a + 1
	fmt.Println("Внутри incrementByValue:", a)
}

func incrementByPointer(a *int) {
	*a = *a + 1
	fmt.Println("Внутри incrementByPointer:", *a)
}

func main() {
	value := 10
	pointer := 10

	fmt.Println("До incrementByValue:", value)
	incrementByValue(value)
	fmt.Println("После incrementByValue:", value)

	fmt.Println("\nДо incrementByPointer:", pointer)
	incrementByPointer(&pointer)
	fmt.Println("После incrementByPointer:", pointer)
}
