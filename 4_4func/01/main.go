package main

import "fmt"

//call by value 기법. 값에 의한 호출
func address(a int) {
	a = a + 3
	fmt.Println(a)
}

func addressRefer(a *int) {
	*a = *a + 3
	fmt.Println(a)
}

func main() {
	a := 3

	addressRefer(&a)
	fmt.Print(&a)
}
