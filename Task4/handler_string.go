package main

import (
	"fmt"
	"strings"
)

func main() {
	var text string = "Olofmeister, I love You"

	fmt.Println("Количество  символов:", len(text), "символов")
	fmt.Println("Есть 'I'?", strings.Contains(text, "I"))
	fmt.Println("Верхний регистр:", strings.ToUpper(text))
}
