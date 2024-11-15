package main

import (
	"fmt"

	"github.com/headfirstgo/keyboard"
)

func main() {
	var gpa [3]float64

	for i := 0; i < len(gpa); i++ {
		fmt.Print("input float num:")
		gpa[i], _ = keyboard.GetFloat() //go get github.com/headfirstgo/keyboard
	}
	for index, value := range gpa {
		fmt.Printf("%d: %f\n", index, value)
	}

}
