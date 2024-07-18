package main

import "fmt"

func fibonacchi(n int) int {

	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}

	return fibonacchi(n-1) + fibonacchi(n-2)

}

var Howmany int = 1

func move_disk(disk_num int, start_peg int, end_peg int) {

	fmt.Printf("%d번 원반을, %d번 기둥에서 %d번 기둥으로 옮겼습니다\n %d번째 이동입니다.\n", disk_num, start_peg, end_peg, Howmany)
	Howmany++
}

func hanoi(N, start_peg, end_peg, otherpeg int) {
	if N == 1 {
		move_disk(1, start_peg, end_peg)
		fmt.Println("여기서 닫힘")
	} else {
		hanoi(N-1, start_peg, otherpeg, end_peg)
		move_disk(N, start_peg, end_peg)
		hanoi(N-1, otherpeg, end_peg, start_peg)
	}

}

func factorial(N int) int {

	if N == 4 {
		return 24
	}

	return N * factorial(N-1)
}

/*
func reverse(s string) string {
	// 기본 사례: 문자열 길이가 0이거나 1인 경우
	if len(s) <= 1 {
		return s
	}
	// 재귀 사례: 첫 문자 + 나머지 문자열의 반전
	return reverse(s[1:]) + string(s[0])
}
*/

//0 1 1 2 3 5
func main() {

	//fmt.Println(reverse("it's too hard"))
	hanoi(3, 1, 3, 2)
	fmt.Println(factorial(5))
}
