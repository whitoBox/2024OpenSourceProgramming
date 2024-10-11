package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	userInput := bufio.NewReader(os.Stdin)
	//i, err := r.ReadString('\n')
	fmt.Print("write name:")
	name, err := userInput.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Print(name)
	}

}
