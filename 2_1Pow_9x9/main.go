package main

import (
	"fmt"
	"math"
)

//제곱구구단

func main() {
	//var i int
	//var j int

	for i := 2; i < 10; i++ {
		fmt.Printf("%d 제곱단", i) //i1 := float64(i)
		for j := 1; j < 10; j++ {
			//j1 := float64(j)
			fmt.Printf("%d 의 %d 제곱  = %d\n", i, j, int(math.Pow(float64(i), float64(j))))
			//fmt.Print(math.Pow10(i)) //10의 n승
		}
		fmt.Println("\n")

	}

}
