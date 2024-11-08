// print prime number between two input numbers
package main

import (
	"fmt"
	"log"

	"main.go/keyboard"
)

func isPrime(num int) bool { //매게변수 / 리턴타임

	if num < 2 {
		return false
	} else if num == 2 {

		return true
	} else if num%2 == 0 {
		return false
	} else {
		for j := 3; j*j <= num; j += 2 { //sqrt(float), int랑 비교하기 위해서 int
			if num%j == 0 {
				return false
			}

		}
	}
	return true
}

func main() {
	fmt.Printf("시작 정수 입력: ")
	n1, err := keyboard.GetInt()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("끝정수 입력: ")
	n2, err := keyboard.GetInt()
	if err != nil {
		log.Fatal(err)
	}

	for i := n1; i <= n2; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}

}
