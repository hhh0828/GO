package main

import (
	"bufio"
	"fmt"
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
	Memo[0] = 0
	Memo[1] = Readint()

	//N 번째 계단에서의 최대값.
	Switching := true
	Code := 0
	for i := 2; i <= N; i++ {
		//내가 선택한 계단,
		Memo[i] = Readint()
		if Switching {
			if Max(Memo[i-1], Memo[i-2]) == Memo[i-2] {
				Switching = false
			}
			Memo[i] += Max(Memo[i-1], Memo[i-2])
			Code++
			if Code == 2 {
				fmt.Println("두칸2번올랐다...")
				Switching = true
			}
		} else {
			fmt.Println("선택 2칸오르키")
			Memo[i] += Memo[i-2]
			Code = 0
			Switching = false
		}

		//세 계단을 연속해서 오르는 경우를 빼줘야함.
		//계단을 하나씩 두번 오른경우는 false로 스위칭.하여

	}
	fmt.Println(Memo[N])

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
