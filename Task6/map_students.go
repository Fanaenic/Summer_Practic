package main

import (
	"fmt"
	"os"
)

func main() {
	StudentList := make(map[string]int)

	StudentList["Макс"] = 5
	StudentList["Петр"] = 5
	StudentList["Александра"] = 4
	StudentList["Олег"] = 5
	StudentList["Ростислав"] = 2
	StudentList["Мстислав"] = 3

	for {
		fmt.Println("1. Показать все оценки группы")
		fmt.Println("2. Добавить оценку")
		fmt.Println("3. Поиск по имени")
		fmt.Println("4. Удаление Студента")
		fmt.Println("5. Выйти из Программы")
		fmt.Println("Выберите действие: ")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("Оценки Текущей Группы\n")
			for name, mark := range StudentList {
				fmt.Println("", name, ":", mark)
			}
		case 2:
			var name string
			var mark int
			fmt.Println("Выберите Имя Студента: ")
			fmt.Scan(&name)
			fmt.Println("Выберите Оценку Студента: ")
			fmt.Scan(&mark)
			StudentList[name] = mark
			fmt.Println("Оценка была успешна занесена в таблицу\n")
		case 3:
			var name string
			fmt.Println("Выберите Имя Студента для полной информации")
			fmt.Scan(&name)
			if mark, exist := StudentList[name]; exist {
				fmt.Println("Оценка Студента:", name, "-", mark)
			} else {
				fmt.Println("Студент Не найден. Попробуйте Заново.")
			}
		case 4:
			var name string
			fmt.Println("Выберите Имя Студента, которого хотите удалить: ")
			fmt.Scan(&name)
			if _, exist := StudentList[name]; exist {
				delete(StudentList, name)
				fmt.Println("Студент был Успешно Удален\n")
			} else {
				fmt.Println("Студент не Найден\n")
			}
		case 5:
			fmt.Println("Выход из программы\n")
			os.Exit(0)
		default:

			fmt.Println("Некорректный выбор, попробуйте снова\n")
		}
	}
}
