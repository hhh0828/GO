package main

//baekjoon dynamic programming memoization
import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	fmt.Print(fibo(n), " ", n-2)

}

var memof [41]int

func fibo(n int) int {
	if n == 2 {
		return 1

	}
	if n == 1 {
		return 1
	}
	if memof[n] == 0 {
		memof[n] = fibo(n-2) + fibo(n-1)
	}
	return memof[n]

}

/*
var memo [40]int
var count int
*/
/*
func fibo2(n int) int {

	memo[0] = 0
	memo[1] = 1
	memo[2] = 1

	for i := 3; i <= n; i++ {
		if memo[i] == 0 {
			memo[i] = memo[i-1] + memo[i-2]
			count++
		}
	}

	return memo[n]
}
*/
// investigate the max value in binary tree..
// need to have a skill that doesn't see the way i have investigated
//Bin value

type Bin struct {
	Left  *Bin
	Right *Bin
	Value int
	Layer int
}

func NewBin(val int, Layer int) *Bin {
	var Bin Bin
	Bin.Layer = Layer
	Bin.Value = val
	return &Bin
}

// WholeBin is below
/*
        7
      3   8
    8   1   0
  2   7   4   4
4   5   2   6   5

0 layer root value

여기서이제 n층까지 가기위한 탐색법으로는 m가지가 있다
m가지의 가짓수크기의 memo를 만들고 만든 메모에서.....
탐색해야한다.








*/

var memo [][]int

func FindMaxbottomup(WholeBin [][]int) int {

	for i := len(WholeBin) - 2; i >= 0; i-- {
		for j := 0; j < len(WholeBin[i]); j++ {
			memo[i][j] += max(memo[i+1][j], memo[i+1][j+1])
		}
	}
	return memo[0][0]
}

func Addelements(WholeBin [][]int) *Bin {
	//root 지정
	root := NewBin(WholeBin[0][0], 0) // 0층의 가장 맨앞 맨앞 값. this is root bin.

	prevlayer := []*Bin{root} //루트설정 시작지설정// prevlayer 초기화
	//level 순회

	for i := 1; i < len(WholeBin); i++ { //2번째 레이어부터 순회
		currentlayer := []*Bin{}                /// 현재 빈 레이어
		for j := 0; j < len(WholeBin[i]); j++ { // 각 레이어 를 순회
			//start from adding the first parents

			bin := NewBin(WholeBin[i][j], i)         // 순회하면서 인자를 bin으로 만들어줌.
			currentlayer = append(currentlayer, bin) // 현재 레이어에 bin을 인자로 넣어줌
			if j == 0 {                              // 첫번째 노드일경우 // 이전 레이어의 j번째 bin의 left값의 bin을 넣어줌
				prevlayer[j].Left = bin
			} else if j == len(WholeBin[i])-1 { // 마지막인덱스일경우
				prevlayer[j-1].Right = bin
			} else {
				prevlayer[j-1].Right = bin
				prevlayer[j].Left = bin
			}

		}
		prevlayer = currentlayer
	}

	return root
}

func Searchbin(tree Bin, n int) int {
	if n == 0 {
		return tree.Value
	}

	return Searchbin(*tree.Right, n-1) + Searchbin(*tree.Right, n-2)
}

//1. add the elements given to the Bintree.
//2.
//Max[N] = Max[N-1] + covalue
//if Max[N] == 0 { add the value }
//전략 1번
/*
make a map [layer][]int
slicing the each line and add the eles to the slice.
add the key/value as layer to map. layer n , layer n-1
and search max in A layer gradually..
to avoid looking into the way that func has looked inside, need to have memoization.


*/

func Searchmax(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {

		}
	}
}
