package main

import "fmt"

// func main() {
// 	a := []int{1, 2, 3, 4, 5}
// 	clone := [5]int{}

// 	for i := 0; i < len(a); i++ {
// 		clone[i] = a[i]
// 	}
// 	fmt.Println(clone)
// }

func main() {
	a := []int{1, 2, 3, 4, 5}
	temp := [5]int{}

	for i := 0; i < len(a); i++ { //i++은 항상고정(i--는 안씀)
		temp[i] = a[(len(a)-1)-i]
		//4-i = 4부터시작 (i=0)

	}
	for i := 0; i < len(a); i++ {
		a[i] = temp[i]
	}
	fmt.Println("temp:", temp)
	fmt.Println("a:", a)
}
