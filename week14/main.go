package main

import "fmt"

func main() {
	ages := make(map[string]int) //맵 자료형 (딕셔너리 같은거)
	//슬라이싱으로 만들어보기 (리스트 같은거)
	var name string
	var age int

	for {
		fmt.Print("what's your name ('q' to exit):")
		fmt.Scanln(&name)
		if name == "q" {
			break
		}
		fmt.Print("what's your age?:")
		fmt.Scanln(&age)

		ages[name] = age
	}
	for name, age := range ages {
		fmt.Printf("%s %d\n", name, age)
	}
}
