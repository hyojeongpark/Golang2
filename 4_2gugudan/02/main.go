package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//변수선언
	var try int
	var score int
	//변수할당
	try = 0
	score = 0
	//시작문구
	print_to_dest() //함수4
	//시도횟수 ++ 맞으면 +1 틀리면 -1 되는 구구단게임만들기
	for {
		n1, n2 := randNumbers() //함수1
		fmt.Printf("%d x %d 는?\n", n1, n2)

		res := n1 * n2
		answer := inputAnswer()       //함수2
		if checkAnswer(answer, res) { //함수3
			score++
			try++
			//순서중요(시도횟수를 더하고 그 값을 출력)
			fmt.Printf("현재점수: %d, 시도횟수: %d\n", score, try)
		} else {
			fmt.Println("틀렸습니다")
			score--
			try++
		}

		if score == 5 {
			fmt.Println("축하합니다! 5점을 얻으셨습니다")
			break
		}
	}
	fmt.Printf("총 시도횟수는 %d번입니다", try)
}

//랜덤숫자를 함수로 만들어보자
func randNumbers() (int, int) { //int를 매개변수로
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(10)
	y := rand.Intn(10)

	return x, y
}

func inputAnswer() int { //반환값 int1개
	fmt.Println("정답은? : ")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	answer, _ := strconv.Atoi(line) //inputAnswer

	return answer

}

func checkAnswer(x, y int) bool { //참 거짓으로 출력
	if x == y {
		return true
	} else {
		return false
	}
}

func print_to_dest() {
	fmt.Println("구구단 게임을 시작하겠습니다!")
	fmt.Println("구구단 게임에서 컴퓨터가 내는 문제를 맞춰 5점을")
	fmt.Println("얻으면 게임은 끝나게 됩니다. (맞출 때 마다 1점씩 증가)")
	fmt.Println("단, 정답이 틀리면 1점 감산됩니다.")
	fmt.Println("=================================")
}
