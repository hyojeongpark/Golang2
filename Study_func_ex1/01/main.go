package main

import "fmt"

func Mist(dur int) {
	step1 := "미스트 뿌리기"
	for i := 1; i <= dur; i++ {
		fmt.Printf(step1+"\t%d"+"회\n", dur)
	}
	// step1 := "미스트 뿌리기"
	// fmt.Printf(step1+"\t%d"+"회\n", dur)
}

//문제: "미스트뿌리기 (dur)회" 가 아닌 숫자만 출력됨
//해결방법: Mist(2) 안에 숫자가 아닌 전체 Mist(2)를 출력하는 함수로 선언하기
//내가사용한방법1: fmt.Printf 11번코드를 반복문안에 넣는다

//1. 매개변수 X, 리턴 X
func Base() {
	fmt.Println("1단계-기초")
	base1 := "skin "
	base2 := "lotion "
	fmt.Println(base1 + "\n" + base2)
}

//2. 매개변수 O, 리턴 X
func Sunscreen(str1 string, str2 string, str3 string, int1 int) {
	fmt.Println("2단계-자외선 차단")
	fmt.Println(str1 + str2 + "/" + str1 + str3)
	fmt.Println(int1)
}

//오류(풀고싶은문제)
//오류내용: unmatched type
//시도 1,2,3...

//3. 매개변수 O, 리턴 O
func FaceMakeup(str1 string, str2 string) string { //string,int같이쓸수는 없는지
	fmt.Println("3단계-피부화장")
	return str1 + "\n" + str2

	//switch문

	// switch str1.(type) {
	// case int:
	// 	fmt.Println("int")
	// case string:
	// 	fmt.Println("string")
	// }
	//문제ㅔ: type이 안맞음??
}

//4. 콜바밸, 콜바레
func EyeMakeup(str1 string, str2 string) {
	fmt.Println("아이메이크업 : ")
	//1. eyebrow 2번 반복하는 반복문 만들기
	//문제1: eyebrow가 아닌 1.a 2.b가 출력됨
	//해결1: str1,str2를 for문 앞으로 배치 (순서가 중요했음)
	//문제2: eyebrow만 2번먼저, eyeliner는 나중에2번 출력하고싶음
	//해결2" for문을 2개 만듬
	str1 = "eyebrow"
	str2 = "eyeliner"
	for i := 1; i < 3; i++ {
		fmt.Println(str1)
	}
	for i := 1; i < 3; i++ {
		fmt.Println(str2)
	}
	//if조건문
	if str1 == "eyebrow" {
		fmt.Printf("성공!\n")
	} else {
		fmt.Printf("실패\n")
	}
}

func ColorMakeup(str1 *string, str2 *string) {
	fmt.Println("립&치크 : ")
	*str1 = "cheeks"
	*str2 = "lip"

	//조건문 if
	if *str1 == "cheeks" {
		fmt.Println(true)
	} else if *str1 == "lip" {
		fmt.Println(false)
	}
	//문제: true,false가 왜 출력이 안되는가?

	fmt.Println(*str1)
	fmt.Println(*str2)
}
func main() {
	fmt.Println("<메이크업 단계>")

	//Mist의 dur(2)만큼 반복
	//횟수 0초기화
	// dur := 0
	// //
	// for i := 0; i < dur+1; i++ {
	// 	//fmt.Println
	// 	Mist(2)
	// }
	//Mist(2) -> 미스트 뿌리기 2회 미스트 뿌리기 2회
	//Mist(3) ->미스트 뿌리기 3회 미스트 뿌리기 3회 미스트 뿌리기 3회
	Mist(3)
	fmt.Println("----------------------------------------------------------")
	Base()

	//Mist(0) 사용안하는 코드는 쓰지않는다
	fmt.Println("----------------------------------------------------------")
	Sunscreen("sun", "screen", "block", 2)

	Mist(3)
	fmt.Println("----------------------------------------------------------")
	a := FaceMakeup("Concealer", "Foundation")
	fmt.Println(a)

	Mist(2)
	fmt.Println("----------------------------------------------------------")
	//기본값
	var str1 string = "1. a"
	var str2 string = "2. b"
	fmt.Println("4단계-색조화장")
	//1. 출력값: eyebrow, eyeliner

	EyeMakeup(str1, str2)
	//2. 출력값: cheeks, lip

	ColorMakeup(&str1, &str2)
	//3. 출력값 : 바뀐값 확인
	fmt.Println("바뀐값 확인 : ")
	fmt.Println(str1, str2)
}
