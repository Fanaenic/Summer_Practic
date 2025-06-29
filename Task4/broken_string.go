package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Смерть стоит того , чтобы жить , а любовь стоит того , чтобы ждать"
	words := strings.Fields(str)
	for i, word := range words {
		fmt.Println(i+1, ":", word)
	}
}
