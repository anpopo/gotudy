package main

import "fmt"

// 변수의 속성 -> 이름, 값, 자료형, 메모리 주소, 사이즈를 가지고 있다.
func main() {
	var a int
	var b int
	var c int
	var d int64
	var flag bool
	var str string
	var ss rune

	a = 3
	b = 4
	c = 7
	d = 99999999999
	flag = true
	str = "안녕 클레오파트라"
	ss = 't'

	fmt.Println(a + b + c)
	fmt.Println(flag)
	fmt.Println(str)
	fmt.Println(d)
	fmt.Println(ss)
}
