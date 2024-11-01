package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("정수 입력: ")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i) //< .strip in python
	num, _ := strconv.Atoi(i)

	isPrime := true

	if num < 2 {
		isPrime = false
	} else if num == 2 {
		isPrime = true
	} else if num%2 == 0 {
		isPrime = false
	} else {
		for j := 3; j*j <= num; j += 2 { //sqrt(float), int랑 비교하기 위해서 int
			if num%j == 0 {
				isPrime = false
				break
			}
			fmt.Printf("%d ", j)
		}
	}

	if isPrime {
		fmt.Printf("%d는 소수입니다.", num)
	} else {
		fmt.Printf("%d는 소수아닙입니다.", num)
	}
}
