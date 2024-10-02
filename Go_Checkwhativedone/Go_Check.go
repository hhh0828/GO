package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
	s := strconv.Itoa(Start)
	e := strconv.Itoa(End)
	//c := strconv.Itoa(Circle)
	Writer.WriteString(s + " " + e + "\n")

}

var Writer *bufio.Writer

func main() {
	Writer = bufio.NewWriterSize(os.Stdout, 200000000)
	//fmt.Println(fibo(9))
	defer Writer.Flush()
	n := 0
	fmt.Scan(&n)

	fmt.Println(hanoi2(n) - 1)
	hanoi(1, 3, 2, n)
}

func hanoi2(n int) int {

	if n == 1 {
		return 2
	}

	return hanoi2(n-1) * 2

}
