package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// a := rand.Intn(50)
	// b := rand.Intn(50)
	// c := rand.Intn(50)
	// d := rand.Intn(50)
	// e := rand.Intn(50)
	for i := 0; i <= 5; i++ {
		ran := rand.Intn(50)
		var sum int
		for j := 1; j <= ran; j++ { //랜덤으로 뽑은 6개의 숫자안에서
			sum = sum + j //총합을 구한다
		}
		fmt.Printf("%d의 1~ %d 사이의 sum(%d)값은 %d이다\n", ran, ran, ran, sum)
	}
	//fmt.Println(a, b, c, d, e)
	//fmt.Println(a + b + c + d + e)

}
