package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	answer := rand.Intn(5) + 5
	//`fmt.Printf("%d\n", answer)

	fmt.Print("input number:")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i) //< .strip in python
	guess, _ := strconv.Atoi(i)

	if answer == guess {
		fmt.Println("정답입니다")
	}
	if answer > guess {
		fmt.Println("값이 작다")
	}
	if answer < guess {
		fmt.Println("값이 크다")
	}
}
