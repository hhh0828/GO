package main

import "fmt"

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

	var data = Data{
		Value: "t",
	}
	data.Array[100] = "t"

	a := ChangeDatafunc2(data)
	fmt.Println(a.Array[100])
	fmt.Println(a.Value)

}

//Pointer의 대한 이해는 좌항은 공간 , 우항은 값 ...

//Why Pointer ?

/*
함수 호출 시. 파라미터에 대한 값 복사(사이즈)가 발생함.
복사된 구조체 또는 변수안에서 해당 함수의 내용이 실행되면서 실제 값이 안바뀌는 일이 발생할 수 있음.
또한 불필요한 메모리사용이 발생 할 수있음, 배열의 사이즈가 큰경우

포인터로 지정하면 8바이트의 주소값으로 같은 값을 가르킬수있다.
예시 below
*/

type Data struct {
	Value string
	Array [200]string
}

// 바로 가르킴 8바이트... 메모리주소값만.. 복사됨
func ChangeDatafunc(arg *Data) {
	arg.Value = "test"
	arg.Array[100] = "test1"
}

// 메모리 낭비
func ChangeDatafunc2(arg Data) Data {
	arg.Value = "test"
	arg.Array[100] = "test1"
	return arg
}

// 궁금점 함수 콜스택과 관련 또한 메모리 사용과 관련하여 가비지컬렉터가 언제 일하냐.
// sweap and mark 마크 하고 지우고 마크하고 지우고...
// 참조가 더이상 발생하지 않을때 제거
