package main

import "fmt"

func main() {
	a := 10
	b := 12
	//산술연산
	c := a + b
	d := a - b
	e := a / b
	f := a * b

	fmt.Println("a+b=", c, d, e, f)

	//논리연산
	fmt.Println("a > b", a > b)
	fmt.Println("a < b", (a < b))
	fmt.Println("a == b", (a == b))
	fmt.Println("a !=b ", (a != b))
	fmt.Println("(a>b) && (a<b)", (a > b) && (a < b))
	fmt.Println("(a>b) || (a<b)", (a > b) || (a < b))

	//비트연산
	fmt.Println("a & b", a&b)
	fmt.Println("a | b", a|b)
	fmt.Println("a ^ b", a^b)
	fmt.Println("a &^ b", a&^b)

	//시프트 연산자
	fmt.Println("a << b ", a<<b)
	fmt.Println("a >> b ", a>>b) // 0000/0000000000101 이라서 앞자리 0만나옴
}
