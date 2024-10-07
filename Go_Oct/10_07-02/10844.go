package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	Reader *bufio.Reader
	Writer *bufio.Writer
)

func init() {
	Reader = bufio.NewReaderSize(os.Stdin, 15000)
	Writer = bufio.NewWriter(os.Stdout)
}

func Readint() int {
	Newintbyte, _, _ := Reader.ReadLine()
	Ni, _ := strconv.Atoi(string(Newintbyte))
	return Ni
}
func main() {

	N := Readint()
	Memo := make([][10]int, N+1)

	for j := 1; j <= 9; j++ {
		Memo[1][j] = 1
	}
	for i := 2; i <= N; i++ {
		for j := 0; j <= 9; j++ {
			if j > 0 {
				Memo[i][j] += Memo[i-1][j-1] // j-1로 끝나는 경우
			}
			if j < 9 {
				Memo[i][j] += Memo[i-1][j+1] // j+1로 끝나는 경우
			}
			Memo[i][j] %= 1000000000
		}
	}
	total := 0
	for j := 0; j <= 9; j++ {

		total += Memo[N][j]
		total %= 1000000000
	}
	//Writer.WriteString(strconv.Itoa(total % 1000000000))
	fmt.Println(total)
	Writer.Flush()

	//fmt.Println(Memo)

	// 45656
	// 인접한 모든수가 1 차이
	// N이주어질때 길이가 N인 계단 수가 총 몇개있는지 구해보자.
	/*
			memo[1] = 9
			memo[2] = 17 > 9 * 2 -1
			memo[3] = memo[2] * 2 - ?
			memo[N] = memo[N-1] * 2 - 1부터N까지의 합


			ex N =1
			1 10 12 / 101 121 123
			2 21 23 / 210 212 232 234
			3
			4
			5
			6
			7
			8
			9 98
			ex N =2
			10 12
			21 23
			32 34
			43 45
			54 56
			65 67
			76 78
			87 89
			98 // 자리수 바뀔때 1빼고
		//0일때는 1로밖에 분열못함
		//9일때는 8로밖에 분열못함.

			   ex N =3
			   101 // 자리수 시작하는 0이있을때 -1 빼고..
			   121 123
			   210 212
			   232 234
			   ...
			   987 989 //(N-1)빼줘야함... 자리수마다

			   N = 4 일경우
			   1212 1234 1232 1210
			   N = 5 일경우
			   12123 12321 12345 12343 12323 12101 12121
			   // 바이트배열로 받아서,
			   // 숫자 체크해서 바이트배열의 인덱스를 돌면서 n-1 또는 n-2 저장
			   // 101 // 0일땐 다음인덱스의 숫자가 1 밖에못옴
			   ,  112 , 110 2의 제곱수로 .., 0일때는 예외처리 .. 해줘야함
				ex N =2
				10 11 20 21 3x9 -1
				11 12 21 23
				32 34 45 43

	*/

}
