package main //기본 준비

import ( //헤더파일이랑 비슷한 개념
	"fmt"
	"reflect"
)

func main() {

	i := 13
	var f float64 = 12.9

	fmt.Printf("value i:%d valuef:%f\n", i, f)
	//fmt.Printf("%d * %f = %f\n", i, f, i*f)//invalid operation: i * f (mismatched types int and float64)
	fmt.Printf("%d * %f = %d\n", i, f, i*int(f))     //conversion
	fmt.Printf("%d * %f = %f\n", i, f, float64(i)*f) //conversion
	fmt.Print(reflect.TypeOf(i), reflect.TypeOf(f))
}
