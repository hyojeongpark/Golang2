//함수 선언식(만드는 방법)
//1. 키워드 함수이름() {}
func Hyo() {
	hyo3 := "jung"
	fmt.Println(hyo3)
	int1 := 1
	int2 := 2
	fmt.Println(int1 + int2)
}

//2. 키워드 함수이름(매개변수 매개변수데이터타입){}
func Hyo1(str string, int1 int, int2 int) {
	hyo3 := str
	fmt.Println(hyo3)
	fmt.Println(int1 + int2)
}

//3. 키워드 함수이름(매개변수 매개변수데이터타입) 데이터타입{}
func Hyo2(num1 int, num2 int) int {
	return num1 + num2
}

//4. 키워드 함수이름() 데이터타입{}  -> 이러한 형태의 함수 만들어 보기 숙제.

//call by value
func change(num1 int, num2 int) {
	num1 -= 1
	num2 -= 2

	fmt.Println(num1)
	fmt.Println(num2)
}

//call by reference
//*int => 포인터로 가리키는 값의 데이터 타입이 int다.
func change1(num1 *int, num2 *int) {
	*num1 -= 1
	*num2 -= 2

	fmt.Println(*num1) //4
	fmt.Println(*num2) //3
}

//실행함수. => go run 명령어 입력시 여기를 찾아서 코드를 일렬씩 수행.
func main() {

	//변수.  변수가 뭘까요? 데이터를 담는 바구니
	//변수 선언하는 방법 1. var 변수명 변수타입
	//var hyo string   //변수 선언 -> 해당 변수를 메모리에 할당한다.
	//hyo = "jung"   //값 할당 -> 해당 변수가 할당된 메모리에 값을 넣어준다.

	//데이터타입 = string, int float, bool, struct, interc
	//var hyo1 string = "jung"

	//변수 선언하는 방법 2. 변수명 := 값할당.
	//hyo2 = "jung"

	//함수 -> 반복되는 코드들을 묶은 코드 블록.
	//함수 호출하는 방법
	//Hyo()
	//Hyo1("jung", 4, 5)

	// a := Hyo2(2, 3) //여기서 2, 3을 뭐라고 할까? -> 인자
	// fmt.Println(a)

	var hyo1 int = 5
	var hyo2 int = 5

	//값이냐 포인터냐 둘중 하나. call by value, call by reference
	//call by value -> 임시적으로 쓸 값을 원본데이터에서 복사해서 인자로 넘김.
	change(hyo1, hyo2)

	//포인터 = 데이터타입. 메모리 상 주소값. 16진수.  *,  &
	fmt.Println("----------------------------------------------------------")
	fmt.Println(hyo1) //5
	fmt.Println(hyo2) //5

	change1(&hyo1, &hyo2)

	fmt.Println("----------------------------------------------------------")
	fmt.Println(hyo1) //4
	fmt.Println(hyo2) //3

}