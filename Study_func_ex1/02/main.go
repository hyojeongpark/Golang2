package main

import "fmt"

func First() {
	step1 := "세수하기"
	dur := 30
	fmt.Printf(step1+"\t%d"+"초\n", dur)
}

func Second(str1 string) {
	step2 := str1
	fmt.Println(step2)
}

func Third(str1 string, str2 string) string {
	return str1 + str2
}

func Fourth(str1 string) {
	str1 = "스크럽"
	fmt.Println(str1)
}
func Fifth(str1 *string) {
	*str1 = "마스크팩"
	fmt.Println(*str1)
}

func main() {
	fmt.Println("스킨케어단계>")
	//1.
	First()

	Second("클렌징 오일")

	a := Third("클렌징", "크림")
	fmt.Println(a)

	var finish string = "수건"

	// //1. 출력값: 수건
	// Fourth(finish)
	// //2. 출력값: 마스크팩
	// Fifth(&finish)
	// //3. 출력값: 마스크팩
	// fmt.Println(finish)

	//1. 출력값: 수건
	Fourth(finish)
	//2. 출력값: 마스크팩
	Fifth(&finish)
	//3. 출력값: 마스크팩
	fmt.Println(finish)
}
