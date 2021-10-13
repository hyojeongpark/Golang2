//함수로 만들어보기
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	money := 3000
	name := "박씨네 과일가게"

	for {
		//규칙적인 패턴일때는 ->함수로 만들어서쓰자
		fmt.Print("무슨 과일을 사실 건가요? \n바나나-1,사과-2,복숭아-3 : \n")
		//1.바나나2.사과3.복숭아나오도록
		num_fruit := inputFruitName()
		//boxes와 switch문의 1,2,3번 일치
		fmt.Printf("고르신 과일은 %s입니다", box)
		box, price := checkname(num_fruit) //※강사님 멘토링) box,price:=
		//나온과일과 갯수를 곱하는 함수
		sum_price := sum(box, price) //※강사님 멘토링) 인자로 box,price
		//
		money = money - sum_price
		//잔액이 부족하면 끝내기
		if money < 0 {
			fmt.Println("잔액부족")
			break
		} else {
			fmt.Printf("남은 소지금액은 %d입니다\n ", money)
			fmt.Printf("총합은 %d입니다\n ", sum_price)
			fmt.Printf("주문하신 가게이름은 %s입니다\n", name)
		}
		//else money < sum_price {
		// 	fmt.Println("소지금이 총합보다 적습니다")
		// 	break
		// }

	}
}
//과일번호 입력함수
func inputFruitName() int {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	num_fruit, _ := strconv.Atoi(line)

	return num_fruit
}
//과일번호별이름 출력함수
func checkname(num_fruit int) (string, int) {
	var box string
	var price int

	switch num_fruit {

	case 1:
		box = "바나나"
		price = 700
	case 2:
		box = "사과"
		price = 1000
	case 3:
		box = "복숭아"
		price = 1500
	default:
		fmt.Println("과일을 잘못입력했습니다")
		break
	}
	return box, price //만든 함수는 반드시 main안에 for문에 넣어야하나?
}
//과일 갯수 입력
func howmany() {
	fmt.Printf("%s는 몇개 사실 건가요?\n", box)
	_, err := fmt.Scanf("%d", &ea)

	if err != nil {
		fmt.Println("잘못입력했습니다")
	} else {
		fmt.Printf("%s는 %d개 샀습니다\n", box, ea) //반환하는 값 유의(주소값아님!)
		break
	}
}

func sum(price int, box string) (int, string) {
	var sum_price int
	var ea int

	sum_price = price * ea
	fmt.Printf("%s의 가격은 총 %d입니다\n", box, sum_price)
	return sum_price
}
