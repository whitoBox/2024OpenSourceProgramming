package main

import (
	"fmt"
	"reflect"
)

func main() {
	gpa := [5]float64{1.1, 2.2, 3.3, 4.4, 4.5}
	var gpaSlice []float64 = gpa[1:4]

	fmt.Println(gpaSlice, reflect.TypeOf(gpaSlice))
	fmt.Println(gpa, reflect.TypeOf(gpa))

}
