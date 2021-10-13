package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//_, c := add(4, 5) //_ =>빈 식별자는 출력하지 않음

	fmt.Println(add(2,3))
	fmt.Println(sub(2,3))
	fmt.Println(div(2,3))
	fmt.Println(mul(2,3))

	num1,num2 :=inputNumber()
	operator := inputOperator()
	result  := calculatorNum(num1,num2,operator)
	printResult(num1,num2,operator,result)


	func add(x, y int) int { //함수명(매개변수),(반환타입)
		return x+y
	}
	func sub(x, y int) int {
		return x-y
	}
	func div(x, y int) int {
		return x/y
	}
	func mul(x, y int) int {
		return x*y
	}
//숫자입력
func inputMessage1(){
	fmt.Println("숫자를 입력하세요: ")
}
//연산자입력
func inputMessage2(){
	fmt.Println("연산자를 입력하세요: ")
}
//숫자넣기
func inputNumber()(int,int){
	inputMessage1() //문구출력
	reader :=bufio.NewReader(os.Stdin)
	line,_:=reader.ReadString('\n')
	line=strings.TrimSpace(line)
	num1,_ :=strconv.Atoi(line)
	inputMessage1()
	reader2:=bufio.NewReader(os.Stdin)
	line2,_:=reader2.ReadString('\n')
	line2=strings.TrimSpace(line2)
	num2,_:=strconv.Atoi(line2)
	
	return num1,num2
}
//연산자 반환
func inputOperator() string{
	inputMessage2()
	reader :=bufio.NewReader(os.Stdin)
	operator,_:=reader.ReadString('\n')
	operator=strings.TrimSpace(operator)

	return operator 
	
}
//숫자 가변인가2개(int), 연산자1개(strint), 연산결과는 int
func calculatorNum(x, y int, o string) int{
	var result int

	switch  o {
	case "+":
		result= x+y
	case "-":
		result= x-y
	case "/":
		result= x/y
	case "*":
		result= x*y
	default: 
		defer fmt.Println("연산자가 아닙니다") 
		//값을 반환하고 출력했을 때 출력(마지막으로 추가)  
	}
	return result
}

func printResult(res int){
	fmt.Println("결과값은 %d 입니다", res)
}
