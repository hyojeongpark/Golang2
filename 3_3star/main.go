package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println() //개행시켜주기
	}
}
