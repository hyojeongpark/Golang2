package main

import "fmt"

func main() {
	var money int
	var menu string

	// 입력함수와 조건문 그리고 반복문을 이용해서 자판기 생성
	for { //<1번 for문>
		fmt.Printf("소지하고 있는 금액 입력 (10000 ~ 20000) : \n")
		//에러변수생성
		_, err := fmt.Scanf("%d\n", &money) //_는 money를 의미함
		if err != nil {
			fmt.Println("잘못 입력했습니다.")
		} else {
			if money < 9999 || money > 20001 {
				//9999원 이하나 20001을 입력하면 무한루프를 돈다
				continue
			} else {
				if money >= 10000 && money <= 20000 {
					break //금액을 잘 입력했으니 <1번 for문>을 빠져나온다
				}
			}
		}
	}
	//for 문과 조건문을 이용해서 자판기를 만들어내는 main 함수를 만들어주세요.
	//무한 반복문 안에 menu 입력하는 함수 scanf를 이용해 메뉴를 입력받고
	// 메뉴의 금액에 따라 소지하고 있는 금액을 줄여 금액이 바닥 날때까지
	// 자판기를 돌려서 금액이 바닥나면 무한 반복문을 탈출하는 프로그램을 만들면 됩니다.
	for { //<2번 for문>
		// 메뉴 입력 받는 곳
		for { //<3번 for문>
			fmt.Print("마실 음료수: \n")

			_, err := fmt.Scanf("%s\n", &menu)

			if err != nil { //에러가 있다
				fmt.Println("잘못 입력했습니다.") //값을 입력하지 않았을 경우
			} else {
				break //메뉴를 잘 입력했을 경우 <3번 for문>을 빠져나온다
			}
			// switch 문으로 case 별로 메뉴와 가격을 정해서 menu 값이 그 메뉴라면 소지금액에서 메뉴 가격을 -- 하면 됨.
			// 소지금액이 바닥나는 조건문은 if문을 통해 소지금액이 0보다 작으면 바깥 무한 반복문을 탈출하는 코드
		}
		switch menu { //<2번 for문>에서 무한반복중
		case "콜라":
			money = money - 2000
			fmt.Printf("남은 소지금액은 %d입니다.", money)
		case "생수":
			money = money - 1000
			fmt.Printf("남은 소지금액은 %d입니다.", money)
		case "게토레이":
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
		}
	}
}
