package main //기본 준비

import ( //헤더파일이랑 비슷한 개념
	"fmt"
	"reflect"
)

func main() {

	i := 13
	var f float64 = 12.9
	c1 := 'A' //룬 4바이트 사용중
	c2 := '김' //https://www.unicode.org/charts/

	fmt.Printf("value i:%d valuef:%f\n", i, f)
	//fmt.Printf("%d * %f = %f\n", i, f, i*f)//invalid operation: i * f (mismatched types int and float64)
	fmt.Printf("%d * %f = %d\n", i, f, i*int(f))     //conversion
	fmt.Printf("%d * %f = %f\n", i, f, float64(i)*f) //conversion

	fmt.Print(c1, c2, "\n")
	fmt.Print(reflect.TypeOf(i), reflect.TypeOf(f), reflect.TypeOf(c1), reflect.TypeOf(c2))

	var ff float64
	var ii int
	var b bool
	var s string

	fmt.Println(ff, ii, b, s)
	fmt.Printf("%f %d %t %s\n", ff, ii, b, s) //zero value
	f = 2.7
	i = 3
	fmt.Print(f > float64(i)) //comparison (true/false)

}
