package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetInt() (int, error) {
	fmt.Print("input number:")
	r := bufio.NewReader(os.Stdin)
	a, err := r.ReadString('\n')
	if err != nil {
		//log.Fatal(err)
		return 0, err
	}

	a = strings.TrimSpace(a)    //< .strip in python
	num, err := strconv.Atoi(a) //<정수로 변환
	if err != nil {
		//log.Fatal(err)
		return 0, err
	}
	return num, nil
}
