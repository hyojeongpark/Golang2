package main

import "fmt"

func main() {
	fmt.Println(f1(10))
}

func f1(a int) int {
	if a == 0 {
		return 1
	} else if a == 1 {
		return 1
	}
	//웬만하면 반환문을 많이 쓰지않는게 좋기때문에
	//피보나치수열일때 아니면 아래처럼 작성

	// if a == 0 || a == 1 {
	// 	return 1
	//  }
	return f1(a-1) + f1(a-2)

}
