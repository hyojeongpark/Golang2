package main

import "fmt"

func main() {
	//10개 공간만들어주기
	var A [10]int

	//10개 공간에 데이터할당해주기
	for i := 0; i < len(A); i++ { //len(A) A의길이
		A[i] = i * i
	}
	fmt.Println(A)
}
