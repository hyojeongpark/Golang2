package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var money int

	// 입력함수와 조건문 그리고 반복문을 이용해서 자판기 생성
	for { //<1번 for문>
		fmt.Printf("소지하고 있는 금액 입력 (10000 ~ 20000) : \n")
		//Scanf보단 빠르지만 공간을 많이 씀
		//bufio=버퍼라는 공간에 데이터를 읽는 객체 생성
		//운영체제 기본 입력 객체 Stdin
		reader := bufio.NewReader(os.Stdin)
		//line변수에 대입
		line, _ := reader.ReadString('\n')
		//눈에 보이지 않는 문자열공백이나 공백 등 쓰레기 값을 없앰
		line = strings.TrimSpace(line)
		//line변수 안 문자열을 컨버터(strconv)로 Int타입으로 변환시켜주기(Atoi)
		n1, _ := strconv.Atoi(line)
		//에러변수생성
		if n1 < 9999 || n1 > 20001 {
			//9999원 이하나 20001을 입력하면 무한루프를 돈다
			continue
		} else {
			money = n1
			break //금액을 잘 입력했으니 <1번 for문>을 빠져나온다
		}
	}
	//for 문과 조건문을 이용해서 자판기를 만들어내는 main 함수를 만들어주세요.
	//무한 반복문 안에 menu 입력하는 함수 scanf를 이용해 메뉴를 입력받고
	// 메뉴의 금액에 따라 소지하고 있는 금액을 줄여 금액이 바닥 날때까지
	// 자판기를 돌려서 금액이 바닥나면 무한 반복문을 탈출하는 프로그램을 만들면 됩니다.
	for { //<2번 for문>
		// 메뉴 설명
		fmt.Println("------------------ -------------------------------------------------------------------------------------")
		fmt.Printf("\t\t\t\t\t\t\t\t음료수 메뉴\n")
		fmt.Printf("   1. 콜라 500원    2. 사이다 700원   3. 생수 600원   4. 게토레이 1200원   5 이상. 무한 나머지음료 1000원 \n")
		fmt.Println("------------------------------------------------------------------------------------------------------")
		fmt.Println("마실 음료수를 고르세요.")
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		menu, _ := strconv.Atoi(line)
		// 메뉴 골랐을 때 각 음료수마다 가격에 따라 소지금액 감산

		switch menu { //<2번 for문>에서 무한반복중
		case 1:
			money = money - 2000
			fmt.Printf("남은 소지금액은 %d입니다.", money)
		case 2:
			money = money - 1000
			fmt.Printf("남은 소지금액은 %d입니다.", money)
		case 3:
			money = money - 3000
			fmt.Printf("남은 소지금액은 %d입니다.", money)
			//case 5개
		default:
			money = money - 5000
			fmt.Printf("남은 소지금액은 %d입니다.", money)
		}
		if money <= 0 {
			fmt.Println("금액이 부족합니다.")
			break //금액이 부족해야 <2번 for문>빠져나올 수 있고 실행이 끝남
		} else {
			fmt.Printf("남은 소지 금액은 %d입니다.\n", money)
		}
	}
}
