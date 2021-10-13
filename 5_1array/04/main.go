package main

import "fmt"

func main() {
	a := make([]int, 0, 3)

	for i := 0; i <= 15; i++ {
		a = append(a, i)
		fmt.Println(len(a), cap(a))
	}
	//fmt.Println(a)
}
