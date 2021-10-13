package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(200)
	b := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)
	// ran1 := rand.Intn(200)
	// ran2 :=rand.Intn(100)
	var count int

	for i := 0; i < a+1; i++ {
		for j := 0; j < a+1; j++ {
			if (i + j) == b { //100이하의 숫자이면
				fmt.Printf("(%d, %d)\n", i, j)
				count++
			}
		}
	}
	fmt.Printf("ran1=%d ", a)
	fmt.Printf("ran2=%d ", b)
	fmt.Printf("총 출력한 갯수는 %d 개", count)
}
