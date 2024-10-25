package main

import (
	"fmt"
)

func main() {
	result := fmt.Sprintf("%v 나누기 %v는 %0.1v 입니다.\n", 1.0, 3.0, 1.0/3.0)
	print(result)
}
