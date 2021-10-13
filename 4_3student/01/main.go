package main

import "fmt"

func main() {
	a, b, c := examGrade(80, 70, 100)
	e, f, g := examSubject("과학", "체육", "미술")
	avg := avg(a, b, c)
	printAVG(avg, e, f, g)
}

//시험과목갯수

func examSubject(r1, r2, r3 string) (string, string, string) {
	return r1, r2, r3

}

func examGrade(o1, o2, o3 int) (int, int, int) {
	return o1, o2, o3
}

//과목과 과목점 수를 가져와 평균을 구하는 함수를 만들어보자. 'func avg(suject [lstring, grade [lint), float64 {
func avg(o1, o2, o3 int) float64 {
	var avg float64
	var sum int

	sum = o1 + o2 + o3
	avg = float64(sum / 3)

	return avg
}
func printAVG(avg float64, a, b, c string) {
	fmt.Printf("평준점수: %f\n", avg)
	fmt.Printf("과목: %s,%s,%s", a, b, c)
}
