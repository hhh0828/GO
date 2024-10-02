package main

//재귀, dynamic program 과 친해져보기...
import "fmt"

func main() {
	//n := 0

	//fmt.Scanln(&n)
	//fmt.Println(factorialrecur(n))

	fmt.Println(factoraldynp(10))
}

//팩토리얼>
/*
N = (N-1) * N
//

// 큰수를 대비 한 스택메모리 절약을 위해
*/

var m [21]int

func factorialrecur(n int) int {

	if n == 0 { // 정보 0! 은 1이었다...
		return 1
	}
	if n == 1 { // n은 1일때 1
		return 1
	}
	if n == 2 { // n은 2일때 2 * 1
		return 2
	}

	if m[n] == 0 {
		m[n] = n * factorialrecur(n-1)
	}
	return m[n]
}

//바텀업 방식이라고 하더라. //n값 20이라고 가정하고...
func factoraldynp(n int) int {

	m[0] = 1
	m[1] = 1
	m[2] = 2
	m[3] = 6

	for i := 4; i < 21; i++ { //연산 2번 O(n)
		m[i] = m[i-1] * i //연산 3번 O(1) //빅오 표기법 헷갈림 다시보기
	}
	return m[n]
}
