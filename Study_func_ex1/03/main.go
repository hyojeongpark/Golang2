package main

import "fmt"

func a(b int) { //출력횟수 b를 매개변수로 넣어주기
	for i := 1; i <= b; i++ { //출력횟수만큼
		fmt.Println(b) //출력하는 내용 :a(안에) b인자만큼 반복
	}
}
func main() {
	a(5) //출력하는 횟수
}
