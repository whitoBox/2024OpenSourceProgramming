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

	for guesses := 0; guesses < 3; guesses++ {
		fmt.Printf("%d 번의 기회 남음 input number:", 3-guesses)

		r := bufio.NewReader(os.Stdin)
		i, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		i = strings.TrimSpace(i) //< .strip in python
		guess, _ := strconv.Atoi(i)

		if answer == guess {
			fmt.Println("correct")
		}
		if answer > guess {
			fmt.Println("your num small")
		}
		if answer < guess {
			fmt.Println("your num big")
		}
	}
}
