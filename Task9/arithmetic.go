package main

import (
	"fmt"
	"math"
)

func main() {

	var num1, num2 float64

	fmt.Print("Энциклопедия по Арифметическим операциям\n")
	fmt.Println("Выберите первое число: ")
	fmt.Scan(&num1)
	fmt.Print("Выберите второе число: ")
	fmt.Scan(&num2)

	fmt.Print("Сложение: ", num1+num2, "\n")
	fmt.Print("Вычитание: ", num1-num2, "\n")
	fmt.Print("Произведение: ", num1*num2, "\n")
	if num2 == 0 {
		fmt.Println("Ошибка! На ноль делить НЕЛЬЗЯ!!!")

	} else {
		fmt.Println("Деление: ", num1/num2, "\n")
	}
	fmt.Println("Возведение в степень: ", math.Pow(num1, num2), "\n")
	fmt.Print("Остаток от деления: ", math.Mod(num1, num2), "\n")
}
