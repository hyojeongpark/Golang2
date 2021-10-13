package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//변수선언
	var cost int
	//변수입력
	cost = inputCost()

	//반복문시작
	for {
		print_menu()
		//메뉴를 번호로 받는 함수호출
		num := inputMenuNumber()
		//switch문으로 남은돈 계산하기
		cost = switchtoCost(num, cost)

		if cost < 0 {
			fmt.Println("잔액부족")
			break
		} else {
			// fmt.Println("남은 소지금액은 %d입니다 ", cost)
		}
	}

}

//함수로 만들기
func inputCost() int { // 매개변수가 왜 필요없을까?사용자에게 받는거라서?
	var cost int //다시 정의해줘야함
	for {
		fmt.Printf("소지하고 있는 금액 입력 (10000 ~ 20000) : \n")
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		cost, _ = strconv.Atoi(line)

		if cost >= 10000 || cost <= 20000 {
			break
		}
	}
	return cost
}

func print_menu() {
	//메뉴출력
	fmt.Println("------------------ -------------------------------------------------------------------------------------")
	fmt.Printf("\t\t\t\t\t\t\t\t아이스크림 메뉴\n")
	fmt.Printf("   1. 설레임 2000원    2. 비비빅 1000원   3. 베스킨 3000원   4. 하겐 5000원   5 이상. 무한 나머지 1000원 \n")
	fmt.Println("------------------------------------------------------------------------------------------------------")
	fmt.Println("아이스크림을 고르세요.")

}

func inputMenuNumber() int {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	num, _ := strconv.Atoi(line)

	return num
}

func switchtoCost(num, cost int) int {
	switch num {
	case 1:
		cost = cost - 2000
		fmt.Printf("남은 소지금액은 %d입니다.\n", cost)
	case 2:
		cost = cost - 1000
		fmt.Printf("남은 소지금액은 %d입니다.\n", cost)
	case 3:
		cost = cost - 3000
		fmt.Printf("남은 소지금액은 %d입니다.\n", cost)
		//case 5개
	default:
		cost = cost - 5000
		fmt.Printf("남은 소지금액은 %d입니다.\n", cost)
	}
	return cost
}
