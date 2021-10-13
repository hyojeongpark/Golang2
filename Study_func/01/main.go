package main

import "fmt"

// func Plus(num1 int, num2 int) int {
// 	return num1 + num2
// }
//call by value
func Change(num1 int, num2 int) {
	num1 -= 1
	num2 -= 2

	fmt.Println(num1)
	fmt.Println(num2)
}

//call by reference
//*int=>포인터로 가리키는 값의 데이터형태가 int다
func Change1(num1 *int, num2 *int) {
	//num1 -= 1 //변수명이고 주소값이기 때문에 int형 연산이 안된다
	//num1은 "포인터 데이터타입"
	*num1 -= 1 //주소값에 담겨있는 값을 찾아감 *num1
	*num2 -= 2

	fmt.Println(*num1) //4
	fmt.Println(*num2) //3
}

func main() {
	//a := Plus(2, 3) //2,3은 "인자"라고 부름, a라는 바구니에 return 값을 넣는다
	//fmt.Println(a)

	var hyo1 int = 5
	var hyo2 int = 5
	//call by value(주소가 바뀌는 것이 아님)
	Change1(&hyo1, &hyo2) //매개변수로 포인터를 받아서 *int
	// *, &의 쓰임

	fmt.Println(hyo1) //4
	fmt.Println(hyo2) //3
}

//
