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
	var aOrNot string

	if score >= 65 {
		aOrNot = "A"
	} else {
		aOrNot = "F"
	}
	fmt.Printf("score:%d %s", score, aOrNot)
}
