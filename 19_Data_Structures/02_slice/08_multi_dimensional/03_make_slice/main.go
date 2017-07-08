package main

import "fmt"

func main () {

	student := make([]string, 3, 20)
	students := make ([][]string, 35)
	//student = append(student, "vinay")
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
}