package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	userInput := bufio.NewReader(os.Stdin)
	//i, err := r.ReadString('\n')
	fmt.Print("write name:")
	name, err := userInput.ReadString('\n')
	fmt.Print(name, err)

}
