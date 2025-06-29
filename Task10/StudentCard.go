package main

import "fmt"

type Student struct {
	Name     string
	Age      int
	Course   int
	AvgGrade float64
}

func createStudent(name string, age, course int, grade float64) Student {
	return Student{
		Name:     name,
		Age:      age,
		Course:   course,
		AvgGrade: grade,
	}
}

func printStudent(s Student) {
	fmt.Printf("Имя: %s\n", s.Name)
	fmt.Printf("Возраст: %d\n", s.Age)
	fmt.Printf("Курс: %d\n", s.Course)
	fmt.Printf("Средний балл: %.2f\n\n", s.AvgGrade)
}

func nextCourse(s *Student) {
	s.Course++
	fmt.Printf("%s переведен на %d курс\n", s.Name, s.Course)
}

func isExcellent(s Student) bool {
	return s.AvgGrade >= 4.5
}

func main() {
	student1 := createStudent("Анна", 20, 2, 4.8)

	printStudent(student1)

	fmt.Printf("%s отличник: %v\n", student1.Name, isExcellent(student1))

	nextCourse(&student1)

	printStudent(student1)
}
