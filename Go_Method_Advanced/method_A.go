package main

import "fmt"

//가변인자 타입Slice타입으로 받음
func sum(nums ...int) int {
	sum := 0
	fmt.Printf("nums 타입은 %T\n", nums)
	for _, v := range nums {
		sum += v
	}
	return sum
}

func main() {
	//OS내의 자원을 반납할때 사용해야함. 파일 관련작업 OS관련작업은 핸들러를 OS로부터 제공받아 사용하기 때문에 함수 종료전에 반납해야함.
	defer fmt.Println("시작은 가장 먼저 시작되나, 지연된 입력입니다.")

	fmt.Println(sum(1, 2, 3, 4))
	fmt.Println(1, 3.14, "ㅋㅋ")
	//아래는 Int 타입만 받음
	//fmt.Println(sum(3.14,2,3,4))

}
