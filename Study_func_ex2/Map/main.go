package main

func main() {
	//key타입=int value타입=string
	//선언 testmap은 nil값. 초기화전
	var testmap map[int]string
	//초기화
	testmap = make(map[int]string)

	testmap[444] = "사과"
	testmap[555] = "배"
	testmap[666] = "포도"

	str1 := testmap[444]
	println(str1)
	println(testmap[444])

	noData := testmap[777]
	println(noData)

	delete(testmap, 444)

	coupon := map[string]string{
		"olive": "young",
		"star":  "bucks",
	}
	a := coupon["olive"]
	println(a)
	// _, exists := coupon["olive"]
	// if exists {
	// 	println("olive yound coupon 15%")
	// }
}
