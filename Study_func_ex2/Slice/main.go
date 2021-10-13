package main

import "fmt"

// func main() {
// 	a := make([]int, 1, 2)
// 	fmt.Println(len(a), cap(a))
// }

// func main() {
// 	a := []string{"바", "보", "야"}
// 	a = a[0:2]
// 	fmt.Println(a)
// }

// func main() {
// 	a := []int{0, 1}
// 	a = append(a, 3) //append안의 a는 문자가 아니라 slice
// 	fmt.Println(a)
// }

func main() {
	a := make([]int, 0, 3)
	for i := 0; i < 5; i++ {
		a = append(a, i)
		fmt.Println(a)
	}

}
