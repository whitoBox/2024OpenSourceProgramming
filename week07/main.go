package main

import (
	"fmt"
	"strings"
)

func main() {
	words := "s#it #appends"
	fix := strings.NewReplacer("#", "h")
	fmt.Print(words + "\n")
	fmt.Print(fix.Replace(words))
}
