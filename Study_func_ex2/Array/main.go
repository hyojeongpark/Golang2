package main

import "fmt"

func main() {
	// var a [3]int //변수명[] 데이터타입
	// a[0] = 1
	// a[1] = 2
	// //a[2] = 3
	// println(a[0])

	//var 변수명 = []데이터타입{데이터들..}
	// var a1 = [3]int{1, 2, 3}
	// var a2 = []int{4, 5, 6, 7, 8}
	// println(a1[1])
	// println(a2[1])

	//배열예제 만들어보기
	//과일가게 배열을 만들고 if=="파인애플"이면 "선택"하기
	var f1 = [3]string{"파인애플", "사과", "포도"}
	fmt.Println(f1)
	if f1[0] == "파인애플" {
		println("선택")
	} else {
		println("다른 코너로 가기")
	}
}
