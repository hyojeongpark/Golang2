package main

import "fmt"

func main() {
	s := "Hello"

	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[i]), ",")
	}
}
