package main

import "fmt"

//This time explains Ftype variable that have got a Func() as its value.

func Mul(a, b int) int {

	return a * b
}

//함수 타입의 별칭 // Int 2개 인자로 받고 Int 반환하는 형
type Opfunction func(int, int) int

//함수 타입을 반환하는것 용도에 따라 다름. 강력함.

//반환 타입은 함수(변수임 ㅋㅋ -_- )
func getOperator(op string) Opfunction {
	if op == "+" {
		return func(a, b int) int {

			return a + b
		}
		//함수 a_b
	} else if op == "*" {
		return func(a, b int) int {

			return a * b
		}
	} else {
		fmt.Println("연산자값 똑바로 입력 바랍니다.")
		return nil
	}
}

//defendency injection - use interface or literal function. those are used to be created when 결합도 줄일때.

func main() {
	oper := getOperator("+")
	oper2 := getOperator("*")
	a := oper(3, 4)
	println(a)
	b := oper2(3, 4)
	println(b)

	c1 := 3
	d1 := 2
	//oper3는 함수타입의 변수 이기때문에, 변할수 있음 레퍼런스 처럼 복사가됨.(포인터가 복사됨)
	oper3 := func(a, b int) int {
		//테스트1 인자로 들어갈 경우에 값이 복사가 됨
		fmt.Println(&a, &b)
		//테스트2 인자로 들어가지 않을 경우 해당함수 내에서 c1,d1을 그대로 포인터로 지목하고 있음
		fmt.Println(&c1, &d1)
		return c1 + d1 + a + b
	}
	println(oper3(c1, d1))

}

//함수 리터럴(람다) 는 내부동작 과정의 상태를 가질 수 있음. first class - inner state.
