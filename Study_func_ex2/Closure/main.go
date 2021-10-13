package main

func closureTest() func() int {
	i := 1
	return func() int {
		for j := 0; j < 3; j++ {
			i++
			//println(i)
		}
		return i
	}
}

func main() {
	keyword := closureTest()

	println(keyword())
	println(keyword())
	println(keyword())
	println("===================")

	//다른이름으로 다시시작
	mykeyword := closureTest()

	for i := 0; i < 7; i++ {
		println(mykeyword())
	}
}
