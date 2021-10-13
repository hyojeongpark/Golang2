package main

import "fmt"

//반복문> for문 : break, continue, goto
//조건문> if문, switch문

//문제> 누적 구구단만들기

func main() {
	//3단 9까지 곱하기
	//3단은 3점 4단은 4점주기
	//sum := 0
	for i := 3; i < 5; i++ {
		for j := 1; j < 10; j++ {
			fmt.Printf("%dX%d=%d\n", i, j, i*j)

		}
		//만약 i=3
		if i == 3 {
			fmt.Printf("점수는 %d점입니다\n", 3)
		} else {
			fmt.Printf("점수는 %d점입니다\n", 4)
		}
	}
}
