package main //기본 준비

import ( //헤더파일이랑 비슷한 개념
	"fmt"
)

func main() {

	var f float64
	var i int
	var b bool
	var s string

	fmt.Println(f, i, b, s)
	fmt.Printf("%f %d %t %s\n", f, i, b, s) //zero value
	f = 2.7
	i = 3
	fmt.Print(f > float64(i)) //comparison (true/false)

}
