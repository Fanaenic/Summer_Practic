package main

import (
	"fmt"
)

func swap(a, b *int) {

	*a, *b = *b, *a

}

func main() {
	x := 33
	y := 44

	fmt.Printf("До изменений: x=%d, y=%d\n", x, y)
	swap(&x, &y)
	fmt.Printf("После изменений: x=%d, y=%d\n", x, y)
}
