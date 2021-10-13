package main

import "fmt"

func main() {
	var a int
	var b int8 //실제로 잘 안씀
	var c int16
	var d int32
	var e int64

	a = -128
	b = 127
	c = 32767
	d = 2000000000
	e = 9223000000000000000

	fmt.Println(a, b, c, d, e)

	//unsigned 양수만 표현
	var f uint
	var g uint8
	var h uint16
	var i uint32
	var j uint64

	f = 999999999999999999
	g = 255
	h = 65535
	i = 4000000000
	j = 9999999999999999999

	fmt.Println(f, g, h, i, j)

	var s string
	s = "hooorrrray"
	fmt.Println(s)

	var r rune
	r = 'R'
	fmt.Println(r)

	var o bool
	o = true
	fmt.Println(o)

}
