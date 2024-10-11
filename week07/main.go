package main

import (
	"fmt"
	"time"
)

func main() {

	var now time.Time = time.Now()
	fmt.Printf("current date: %d.%s.%d\n", now.Year(), now.Month(), now.Day())
	fmt.Printf("current time: %d:%d.%d\n", now.Hour(), now.Minute(), now.Second())
	fmt.Print(now.Month())
}
