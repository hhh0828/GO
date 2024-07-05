package main

import "fmt"

type test struct {
	a int
	b [10]int
}

//구조체에 포인터를 입력을 해준다.... test의 포인터로 가라
func changedata(arg *test) {
	arg.a = 10
	arg.b[9] = 1
	fmt.Println(arg.a, arg.b)

}

//Go, can control pointer like c/c++ but can't use operator for operating.
func main() {
	//var a int
	//int타입에 해당하는 주소값을 가지는 p
	//var p *int
	//a = 8
	// a의 주소값을 대입 a의값은 8
	//p = &a
	//p의 실제 주소값을 가르키면서 20을 대입.
	//pointer의 기본값은 Nil
	//*p = 20
	//fmt.Println(p)
	//fmt.Printf("%d", a)

	var test1 test = test{
		a: 0,
	}
	//주소값을 넣으면
	changedata(&test1)

	fmt.Println(test1.a, test1.b[9])
}

//Pointer의 대한 이해는 좌항은 공간 , 우항은 값 ...
