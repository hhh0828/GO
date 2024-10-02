package main

//재귀, dynamic program 과 친해져보기...
import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	//n := 0

	//fmt.Scanln(&n)
	//fmt.Println(factorialrecur(n))

	//kora := "ABBA"
	//bytekora := []byte(kora)
	//fmt.Println(prh(bytekora, 0, len(bytekora)-1))
	//fmt.Println(pr(bytekora, 0, len(bytekora)-1))
	//fmt.Println(factoraldynp(10))
	T := Getint()
	defer Writer.Flush()
	count = 1 // 들어가는 시점을 1로 ^_^ 0부터시작이아니니까..
	for i := 0; i < T; i++ {
		thestr := Getstrline()

		result := pr([]byte(thestr), 0, len(thestr)-1)
		strconv.Itoa(result)
		tra := strconv.Itoa(count)
		Writer.WriteString(strconv.Itoa(result) + " " + tra + "\n")
		count = 1
	}

}

//편법으로 풀어보기 - 안먹힘 ㅋㅋ - 예외처리안됨.
//정석으로 풀어보자
//펠린드롬의 문자열의 갯수 N/2 +2 >> 홀수 개의 스택수
// N/2 +1 >> 짝수개의 스택수

var (
	Reader *bufio.Reader
	Writer *bufio.Writer
)

func init() {
	Reader = bufio.NewReaderSize(os.Stdin, 10)
	Writer = bufio.NewWriter(os.Stdout)
}

func Getint() int {
	a, _, _ := Reader.ReadLine()
	transferreda, _ := strconv.Atoi(string(a))
	return transferreda
}
func Getstrline() string {
	a, _, _ := Reader.ReadLine()
	return string(a)
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

// 바텀업 방식이라고 하더라. //n값 20이라고 가정하고...
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
