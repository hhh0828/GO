package main

import "fmt"

//import "fmt"

//Bit Shift Overflow test x GO
func main() {

	//var x int8 = 8
	//var y int8 = 64
	// The Go Playground로 이동할 수 있음.

	//x는 8칸 바이너리 빈칸0 대입, x<<2 왼쪽으로 두칸 빈칸 0으로 채운 8칸짜리 바이너리, x<<2의 변화는? 8에서 32가 되었음
	//fmt.Printf("x:%08b x<<2:%08b x<<2: %d\n", x, x<<2, x<<2)
	//y는 8칸 바이너리 빈칸0 대입, y<<2 왼쪽으로 두칸 빈칸 0으로 채운 8칸짜리 바이너리, y<<2의 변화는 64에서 0 >> overflow 가 발생함.
	//fmt.Printf("y:%08b y<<2:%08b y<<2: %d\n", y, y<<2, y<<2)

	//함수 스터디
	//test_Minus_Plus()
	//var a int = 126
	//var b int = 3000
	//var d float64 = 9999
	//var f float64 = 7890

	//다른 go파일에 쓰인 파일은 반드시 저장 후 사용 하기 바람.
	//c := add(a, b)
	//fmt.Println(c)
	//fmt.Println(add(a, b))
	//fmt.Println(Divide(d, f))

	//선언과 동시에 변수 대입하기
	//k, succ := Divide(d, f)
	//fmt.Println(k, succ)

	var a int = 10
	Recursivetest(a)
	fmt.Println("the test has been completed")
}
