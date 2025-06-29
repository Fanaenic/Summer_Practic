package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Добро пожаловать в калькулятор!")
	fmt.Println("Выберите дальнейшие действия в калькуляторе:")

	for {
		fmt.Println("1. Сумма")
		fmt.Println("2. Произведение")
		fmt.Println("3. Деление")
		fmt.Println("4. Разность")
		fmt.Println("5. Корень")
		fmt.Println("6. Квадрат")
		fmt.Println("0. Выход")

		var choice int
		fmt.Print("Выберите операцию: ")
		fmt.Scan(&choice)

		switch choice {
		case 0:
			fmt.Println("Выход из калькулятора.")
			return
		case 1:
			var a, b float64
			fmt.Print("Введите число a: ")
			fmt.Scan(&a)
			fmt.Print("Введите число b: ")
			fmt.Scan(&b)
			fmt.Printf("Результат: %.2f\n\n", a+b)
		case 2:
			var a, b float64
			fmt.Print("Введите число a: ")
			fmt.Scan(&a)
			fmt.Print("Введите число b: ")
			fmt.Scan(&b)
			fmt.Printf("Результат: %.2f\n\n", a*b)
		case 3:
			var a, b float64
			fmt.Print("Введите число a: ")
			fmt.Scan(&a)
			fmt.Print("Введите число b: ")
			fmt.Scan(&b)
			if b == 0 {
				fmt.Println("Ошибка: деление на ноль!\n")
			} else {
				fmt.Printf("Результат: %.2f\n\n", a/b)
			}
		case 4:
			var a, b float64
			fmt.Print("Введите число a: ")
			fmt.Scan(&a)
			fmt.Print("Введите число b: ")
			fmt.Scan(&b)
			fmt.Printf("Результат: %.2f\n\n", a-b)
		case 5:
			var a float64
			fmt.Print("Введите число: ")
			fmt.Scan(&a)
			if a < 0 {
				fmt.Println("Ошибка: нельзя извлечь корень из отрицательного числа!\n")
			} else {
				fmt.Printf("Результат: %.2f\n\n", math.Sqrt(a))
			}
		case 6:
			var a float64
			fmt.Print("Введите число: ")
			fmt.Scan(&a)
			fmt.Printf("Результат: %.2f\n\n", a*a)
		default:
			fmt.Println("Неверный выбор, попробуйте снова.\n")
		}
	}
}
