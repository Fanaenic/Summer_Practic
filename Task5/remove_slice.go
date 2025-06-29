package main

import "fmt"

func removeElement(slice []string, index int) []string {

	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	fruits := []string{"Apple", "Banana", "Orange"}
	fmt.Println("До Удаления:, ", fruits)
	fruits = removeElement(fruits, 1)

	fmt.Println("После Удаления:", fruits)
}
