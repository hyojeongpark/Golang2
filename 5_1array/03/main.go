package main

import "fmt"

// func main() {
// 	s1 := "Hello 월드"

// 	for i := 0; i < len(s1); i++ {
// 		fmt.Print(s1[i], ", ")
// 	}
// }

func main() {
	s1 := "Hello 월드"
	//byte=1 rune=4 string=16바이트
	s2 := []rune(s1)

	for i := 0; i < len(s2); i++ {
		fmt.Print(s2[i], ", ")
	}
}

//월드" 부분은 3개의 값이 하나의 공간에 들어가있는 형태
