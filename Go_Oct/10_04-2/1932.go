package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	Reader *bufio.Reader
	Writer *bufio.Writer
)

func init() {
	Reader = bufio.NewReaderSize(os.Stdin, 10240)
	Writer = bufio.NewWriter(os.Stdout)
}

func Readint() int {
	a, _, _ := Reader.ReadLine()
	tra, _ := strconv.Atoi(string(a))

	return tra
}

func Readline() *[]int {
	a, _, _ := Reader.ReadLine()
	inputs := strings.Split(string(a), " ")
	numarr := make([]int, len(inputs))
	for i := 0; i < len(inputs); i++ {
		cc, _ := strconv.Atoi(inputs[i])
		numarr[i] = cc
	}
	return &numarr
}

func main() {
	defer Writer.Flush()
	N := Readint()
	if N == 1 {
		a := Readint()
		Writer.WriteString(strconv.Itoa(a))
	} else {
		superarr := make([]([]int), 0)
		for i := 0; i < N; i++ {
			superarr = append(superarr, *Readline())
			//fmt.Println(superarr)
		}
		a := Findmax(superarr)

		Writer.WriteString(strconv.Itoa(a))
	}

}

func Findmax(arr2 [][]int) int {

	for i := len(arr2) - 2; i >= 0; i-- {
		for j := 0; j < len(arr2[i]); j++ {
			arr2[i][j] += max(arr2[i+1][j], arr2[i+1][j+1])
			//fmt.Println(arr2[i][j])
		}
	}
	return arr2[0][0]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
