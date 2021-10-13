package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	flag := true

	for flag {
		a := int(rand.Float64()*100 + 100)
		if (a%2)|(a%9) == 0 {
			fmt.Println("랜덤숫자는 2또는 9의 배수입니다", a)
			flag = true
			break

		}

	}

}
