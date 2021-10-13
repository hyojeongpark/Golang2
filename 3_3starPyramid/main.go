package main

import "fmt"

func main() {
	//피라미드
	for i := 0; i < 5; i++ { //1,2,3,4행 이라고 생각하자
		for j := 0; j < 4-i; j++ { //거꾸로 star했던 부분을 공백으로 처리
			fmt.Print(" ")
		}
		for j := 0; j < 2*i+1; j++ { //2번째줄에3,3번째줄에5..공식발견 2*i+1
			fmt.Print("*")
		}
		fmt.Println() //개행시켜주기
	}
	fmt.Println() //개행시켜주기
	//다이아몬드
	for i := 0; i < 3; i++ {
		for j := 0; j < 2-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < i*2+1; k++ {
			fmt.Print("*") //print항목만 반대로?
		}
		fmt.Println()
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print(" ")
		}

		for k := 0; k < 3-i*2; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
