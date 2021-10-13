package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("i값은 : %d\n", i)
	}
	//무한 반복문
	a := 0
	for a == 12 {
		break //반복을 끝냄??
	}
	a++

}
