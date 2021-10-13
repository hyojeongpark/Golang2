package main

import "fmt"

func main() {
	//변수를 선언. 소지금액을 선언 대입문으로 초기화
	money := 3000
	//가게 이름 변수선언
	name := "박씨네 과일가게"

	var num_fruit int //과일번호
	var ea int        //갯수
	var box string
	var price int
	var sum_price int

	for { //for 무한반복
		for {
			//과일 가게메뉴 번호+갯수 입력받기
			//1,2,3번중 입력하도록 만든다
			fmt.Printf("무슨 과일을 사실 건가요? : \n")
			_, err01 := fmt.Scanf("%d\n", &num_fruit)

			if err01 != nil {
				fmt.Println("잘못 입력했습니다.")
			} else {
				break
			}
		}
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
		}

		for {
			fmt.Printf("%s는 몇개 사실 건가요?\n", box)
			_, err := fmt.Scanf("%d", &ea)

			if err != nil {
				fmt.Println("잘못입력했습니다")
			} else {
				fmt.Printf("%s는 %d개 샀습니다\n", box, ea) //반환하는 값 유의(주소값아님!)
				break
			}
		}
		//sum_price 초기화 먼저해주기
		sum_price = price * ea
		fmt.Printf("%s의 가격은 총 %d입니다\n", box, sum_price)

		//무한반복문 = 과일갯수를 입력했을 때 소지금이 모자라면 e탈출
		if money < sum_price {
			fmt.Println("소지금이 총합보다 적습니다")
			break
		} else if money == 0 || money < 0 {
			fmt.Println("잔액부족")
			break
		}
		fmt.Printf("-----주문하신 가게 이름: %s-----", name)
		break
	}
}
