package main

import (
	"bufio"
	"os"
	"strconv"
)

//step up 1 or 2 . and make points

var (
	Reader *bufio.Reader
	Writer *bufio.Writer
)

func init() {
	Reader = bufio.NewReaderSize(os.Stdin, 15000)
	Writer = bufio.NewWriter(os.Stdout)
}

func main() {

	//nums of stairs that needs to be added
	N := Readint()
	Memo := make([]int, N+1)

	Stairs := make([]int, N+1)

	//N 번째 계단에서의 최대값.
	for i := 1; i <= N; i++ {
		// init each stair scores
		Stairs[i] = Readint()
	}
	//fmt.Println(Memo)
	//Dp start
	Memo[0] = 0

	Memo[1] = Stairs[1]
	if N >= 2 {
		Memo[2] = Stairs[1] + Stairs[2]
	}
	for i := 3; i <= N; i++ { // 계단을 오르는 경우의수, 즉 n-1 n-1 n-1 을 제외한.
		//N-2 + current score, N-2 + N-1 + N-1
		Memo[i] += Max(Memo[i-2]+Stairs[i], Memo[i-3]+Stairs[i-1]+Stairs[i])
	}
	//
	Writer.WriteString(strconv.Itoa(Memo[N]))
	Writer.Flush()
}

/*

계단 선택.. N-1 / N-1 / N-1 안됨
*/

func Readint() int {
	Newintbyte, _, _ := Reader.ReadLine()
	Ni, _ := strconv.Atoi(string(Newintbyte))
	return Ni
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
