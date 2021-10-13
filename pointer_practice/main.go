package main

import "fmt"

func funcA(y int) {
	y -= 10        //2. y = y-10 이고, 3-10 = -7이다
	fmt.Println(y) //답 -7
}
func funcB(y *int) {
	*y -= 10 //5. 함수안에서 답은 동일하게 -7 이다
	fmt.Println(y)
}
func main() {
	var a int
	var b int

	a = 3
	b = 3

	funcA(a)       //1. a에 할당한 3을 넣어보자
	fmt.Println(a) //3. 원래 값인 3 출력!
	funcB(&b)      //4. a의 주소값(포인터)을 함수에 넣어보자 (일단 주소값 출력)
	fmt.Println(b) //6. 하지만 여기서도 -7이 된다는게 다른점이다
}
