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
	fmt.Print("score:")
	score := bufio.NewReader(os.Stdin)
	grade, err := score.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	grade = strings.TrimSpace(grade)
	finalGrade, _ := strconv.ParseInt(grade, 10, 32)

	if finalGrade >= 90 {
		fmt.Println("you pass")
	} else {
		fmt.Println("you fail")
	}
}
