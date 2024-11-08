// print prime number between two input numbers
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

func getInt() (int, error) {
	r := bufio.NewReader(os.Stdin)
	a, err := r.ReadString('\n')
	if err != nil {
		//log.Fatal(err)
		return 0, err
	}

	a = strings.TrimSpace(a)    //< .strip in python
	num, err := strconv.Atoi(a) //<정수로 변환
	if err != nil {
		//log.Fatal(err)
		return 0, err
	}
	return num, nil
}

func main() {
	fmt.Printf("시작 정수 입력: ")
	n1, err := getInt()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("끝정수 입력: ")
	n2, err := getInt()
	if err != nil {
		log.Fatal(err)
	}

	for i := n1; i <= n2; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}

}
