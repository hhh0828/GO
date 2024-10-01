package main

import (
	"bufio"
	"os"
	"strconv"
)

//문제에 대한 접근 방법 고찰.
//그 전에.. 이건이 문제인가? 에대한 정의? 인식?
/*
창의성과 지식의 관계
새로운 이론과 가설에 대한 결과가 맞다면 그건 우연의 산물인가 노력의 결과인가 발견인가
셀 수 없이 많은 실패에서 나오는 결과뿐인건가
인사이트를 가지려면 어떻게 해야하는것인가.

//
making segments of probs

why the problem happens?
what is purpose to resolve the problem?





문제에 대한 일반화와 추상화 도식화 절차, 관계 맵핑

*/
//N개의 계단을 1개 또는 2개를 오를수있다고 가정한다.
//그러면 N개의 계단이 주어졌을때 이를 오를 수있는 방법 경우의 수에 대해 생각해보자.
//계단을 1개오르면 N-1이고 계단을 2개오르면 N-2개가 남는다.
//N개의 계단을 오르는 방법의 수는 처음부터 1개의 계단을 오를때의 경우의수 + 2개의 계단을 오를떄 경우의 수 이다...

//일반화 하면 N의 경우의수 = N-1경우의수 + N-2 경우의수.

//재귀의핵심은 문제의 일반화와 마지막값의 결과 일반화값.
//수를 구할때의 일반화는
//재귀로 만들면 가지의 끝부분에서 계단이 2개남았을때 경우의수 2개
//계단이 1개 남았을때 경우의수 1개 를 리턴해주자.

//계단이 0개일때는 0을 리턴.

//

// 문제
var (
	reader *bufio.Reader
	writer *bufio.Writer
)

func init() {
	reader = bufio.NewReaderSize(os.Stdin, 1000)
	writer = bufio.NewWriter(os.Stdout)
}

func main() {
	defer writer.Flush()

	T1, _, _ := reader.ReadLine()

	T, _ := strconv.Atoi(string(T1))

	//fmt.Println(fibo(n, &a))

	for i := 0; i < T; i++ {
		n1, _, _ := reader.ReadLine()
		n, _ := strconv.Atoi(string(n1))
		a := make([]int, n+1)
		printprint(fibo(n-1, &a), fibo(n, &a))
	}

	/*
		a := Stair(100)
		a2 := Stairver2(100)
		fmt.Println(a)
		fmt.Println(a2)
	*/
}

// 입력값 받고 make 함수로..배열초기화 가능.
var returns [101]int

func Stair(N int) int {

	if N == 2 {
		return 2
	}
	if N == 1 {
		return 1
	}
	///이렇게하면 성능이 좋아지지 않을까 기대했지만 택도없음
	if N == 4 {
		return 5
	}
	if N == 3 {
		return 3
	}

	//값이 비어있을떄만 계산하면된다.
	//이미 봤던값을은 계산 안하면 그만임.
	if returns[N] == 0 {
		returns[N] = Stair(N-1) + Stair(N-2)
	}

	return returns[N]
}

// 뭔가 밑에서 올라감 재귀랑 다른맛이있네
func Stairver2(N int) int {
	//흠 계단값이 주어짐.
	a := make([]int, N+1)
	a[0] = 0
	a[1] = 1
	a[2] = 2

	for i := 3; i <= N; i++ {
		if a[i] == 0 {
			a[i] = a[i-1] + a[i-2] //  배열에 랜덤 요소 접근은 O(1)
		}
	}
	return a[N]
	//시간복잡도는 For문한번 순회 O(n)
}

func fibo(n int, m *[]int) int {

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if (*m)[n] == 0 {
		(*m)[n] = fibo(n-2, m) + fibo(n-1, m)
	}
	return (*m)[n]
}

func printprint(a, b int) {
	converteda := strconv.Itoa(a)
	convertedb := strconv.Itoa(b)

	writer.WriteString(converteda + " " + convertedb + "\n")
}
