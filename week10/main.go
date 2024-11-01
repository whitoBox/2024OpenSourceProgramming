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

	//count := 0
	isPrime := true //가독성, 저장공간크기 개선 (변수에 'is-'들어가는건 true/false 들어간다.)
	for j := 2; j < num; j++ {
		if num%j == 0 {
			//count++
			isPrime = false
		}
	}
	if isPrime {
		fmt.Printf("%d는 소수입니다.", num)
	} else {
		fmt.Printf("%d는 소수아닙입니다.", num)
	}
}
