package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Я люблю Мтуси Я люблю книги Я люблю золото"

	words := strings.Fields(text)

	count := make(map[string]int)

	for _, word := range words {
		count[word]++
	}

	fmt.Println("Результат подсчёта:")
	for word, num := range count {
		fmt.Printf("%s: %d\n", word, num)
	}
}
