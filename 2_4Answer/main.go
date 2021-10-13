package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	for {
		none := false
		a := int((rand.New(rand.NewSource(time.Now().UnixNano())).Float64() * 100) + 100)
		//*100 은 100부터라는 의미 +100까지하니 (100~200사이 숫자라는 의미)
		for i := 2; i < 10; i++ {
			if a%i == 0 {
				none = false
				fmt.Printf("%d는 %d의 배수입니다\n", a, i)
				break
			} else {
				none = true
			}
		}
		if none == true {
			fmt.Printf("%d는 2~9사이의 어느 배수도 아니다\n", a)
		} else { //none=false
			break
		}
	}
}
