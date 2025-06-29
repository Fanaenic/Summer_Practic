package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name      string
	BirthYear int
	Course    int
	AvgGrade  float64
}

func createStudent(name string, birthYear, course int, grade float64) Student {
	return Student{
		Name:      name,
		BirthYear: birthYear,
		Course:    course,
		AvgGrade:  grade,
	}
}

func printStudent(s Student) {
	fmt.Printf("Имя: %s\n", s.Name)
	fmt.Printf("Возраст: %d\n", s.CalculateAge())
	fmt.Printf("Курс: %d\n", s.Course)
	fmt.Printf("Средний балл: %.2f\n", s.AvgGrade)
	fmt.Printf("Статус: %s\n\n", s.GetStatus())
}

func nextCourse(s *Student) {
	s.Course++
	fmt.Printf("%s переведен на %d курс\n", s.Name, s.Course)
}

func isExcellent(s Student) bool {
	return s.AvgGrade >= 4.5
}

func (s Student) CalculateAge() int {
	currentYear := time.Now().Year()
	return currentYear - s.BirthYear
}

func (s Student) GetStatus() string {
	switch {
	case s.AvgGrade >= 4.5:
		return "отличник"
	case s.AvgGrade >= 3.5:
		return "хорошист"
	default:
		return "троечник"
	}
}

func main() {
	student1 := createStudent("Анна", 2003, 2, 3)

	printStudent(student1)

	fmt.Printf("%s отличник: %v\n", student1.Name, isExcellent(student1))

	nextCourse(&student1)

	printStudent(student1)
}
