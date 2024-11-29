package main

import "fmt"

func main() {
	var student1 struct {
		id   int
		name string
		gpa  float32
	}
	student1.id = 202444029
	student1.name = "John"
	student1.gpa = 3.98

	fmt.Println(student1)

	var student2 struct {
		id   int
		name string
		gpa  float32
	}
	student2.id = 202444030
	student2.name = "Wick"
	student2.gpa = 4.1
	fmt.Println(student2.id)

}
