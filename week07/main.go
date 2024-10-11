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
	fmt.Print("input score:")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i)                //< .strip in python
	score, _ := strconv.ParseInt(i, 16, 32) //(문자열 , 진수 , 비트)

	if score >= 65 {
		fmt.Printf("score:%d A", score)
	} else {
		fmt.Printf("score:%d F", score)
	}
}
