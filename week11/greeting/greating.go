package greeting

import (
	"fmt"
)

func hi(name string) {
	fmt.Printf("hi %s\n", name)
}

func hello(name string) {
	fmt.Printf("hello %s\n", name)
}

func Allgreetings(name string) {
	hi(name)
	hello(name)
}
