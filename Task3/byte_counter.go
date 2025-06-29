package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var (
		a  int   = 42
		a2 int8  = 42
		a3 int16 = 42
		a4 int32 = 42
		a5 int64 = 42

		b  float32 = 3.14
		b1 float64 = 3.14

		c  complex64  = 1 + 2i
		c1 complex128 = 1 + 2i

		s string = "Oleg"
	)
	fmt.Println("Размер байтов в int: ", unsafe.Sizeof(a), "байт")
	fmt.Println("Размер байтов в int8: ", unsafe.Sizeof(a2), "байт")
	fmt.Println("Размер байтов в int16: ", unsafe.Sizeof(a3), "байт")
	fmt.Println("Размер байтов в int32: ", unsafe.Sizeof(a4), "байт")
	fmt.Println("Размер байтов в int64: ", unsafe.Sizeof(a5), "байт")
	fmt.Println("Размер байтов в float32: ", unsafe.Sizeof(b), "байт")
	fmt.Println("Размер байтов в float64: ", unsafe.Sizeof(b1), "байт")
	fmt.Println("Размер байтов в complex64: ", unsafe.Sizeof(c), "байт")
	fmt.Println("Размер байтов в complex128: ", unsafe.Sizeof(c1), "байт")
	fmt.Println("Размер байтов в string: ", unsafe.Sizeof(s), "байт")
}
