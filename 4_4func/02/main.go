package main

import "fmt"

func closer(x int) func() int {
	//클로저
	return func() int {
		i := func() int {
			se := 3
			return se + 1
		}
		return func() int {
			a := 3 + x
			a += i()
			return a
		}

	}
}
func main() {
	a := closer(3)
	fmt.Println(a()) //답:10
}
