package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ { //1,2,3,4,5
		for j := 0; j < 5-i; j++ {
			fmt.Print("*")
		}
		fmt.Println() //개행시켜주기
	}
}
