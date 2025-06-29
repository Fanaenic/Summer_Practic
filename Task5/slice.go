package main

import "fmt"

func main() {
	var fruits []string

	fruits = append(fruits, "Apple")
	fruits = append(fruits, "Banana", "Orange")

	for i, fruit := range fruits {
		fmt.Printf("%d: %s\n", i+1, fruit)
	}
}
