package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	fmt.Printf("정수 입력: ")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i) //< .strip in python
	num, _ := strconv.Atoi(i)

	if isPrime(num) {
		fmt.Printf("%d는 소수입니다.", num)
	} else {
		fmt.Printf("%d는 소수아닙입니다.", num)
	}
}
