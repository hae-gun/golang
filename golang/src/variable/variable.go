package main

var a int
var f float32 = 11.

func main() {
	a = 10
	f = 12.0
	// go 에서는 선언된 변수가 사용되지 않으면 에러를 발생시킨다.

	var i, j, k int
	i = 1
	j = 2
	k = 3
	println(i + j + k)

	// 함수 안에서 아래 두개는 같은의미 func 내부에서만 사용 가능하다.
	// var 변수명 = value
	// 변수명 := value
	var b = 1
	c := 1

	println(b)
	println(c)

	s := "String~"
	println(s)

	// const 키워드를 통해 상수 생성가능.
	const word = "Word"

	// 묶음으로 선언가능
	const (
		A = "1"
		B = "1"
		C = "1"
	)

	// 상수 0부터 순차적으로 부여할 때 iota identifier 를 사용한다.
	const (
		aa = iota
		bb
		cc
		dd
	)

	println(aa)
	println(bb)
	println(cc)
	println(dd)

}
