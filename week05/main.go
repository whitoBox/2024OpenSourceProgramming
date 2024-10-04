package main //기본 준비

import ( //헤더파일이랑 비슷한 개념
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	//var a int = 10

	var f float32 = 4.3
	g := 4.3 //float64
	i := 10  //옆에 값을 해석하여 자동으로 변수를 만들어줌

	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(g), reflect.TypeOf(i))
	fmt.Printf("%s\n", strings.Title("inha tech"))
	fmt.Println(math.Ceil(3.99), math.Floor(3.4))

	fmt.Printf("value i:%d\n", i)
	fmt.Println("value i:", i)
	fmt.Print("value i:", i)
}
