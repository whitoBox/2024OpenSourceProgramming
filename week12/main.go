package main

//go mod init
import (
	"fmt"
	"time"
)

func main() {
	// var dates [3]time.Time
	// things := [2]int{10, 20}
	// dates[0] = time.Unix(1447920000, 0)

	// fmt.Printf("%d", things)
	//var dates [3]time.Time = [3]time.Time{time.Unix(1, 0), time.Unix(2, 0), time.Unix(3, 0)}

	// dates := [3]time.Time{time.Unix(1, 0), time.Unix(2, 0), time.Unix(3, 0)}
	// fmt.Println(dates[0], dates[1], dates[2])

	dates := [3]time.Time{
		time.Unix(1, 0),
		time.Unix(2, 0),
		time.Unix(3, 0), //줄 바꿀때도 콤마가 필요하다.
	}
	fmt.Println(dates[0], dates[1], dates[2])
}
