package main

import "fmt"

func fibo(N int) int {
	if N < 2 {
		return 1
	} else if N == 0 {
		return 0
	}
	return fibo(N-1) + fibo(N-2)
}

func hanoi(Start, End, Extra, Circle int) {
	if Circle == 1 {
		MoveCircle(Start, End, Circle)
		return
	}
	//N-1를 시작지점에서 보조기둥 Extra로
	//N-1를 다시 끝지점으로 옮겨야함.
	hanoi(Start, Extra, End, Circle-1)
	MoveCircle(Start, End, Circle)
	hanoi(Extra, End, Start, Circle-1)

}

func MoveCircle(Start, End, Circle int) {

	fmt.Printf("%d번째 원판을 %d번째 기둥에서 %d번째 기둥으로 옮겼습니다\n", Circle, Start, End)

}

func main() {

	fmt.Println(fibo(9))
	hanoi(1, 3, 2, 3)
}
