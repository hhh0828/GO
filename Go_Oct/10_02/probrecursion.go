package main

//문제요약 나선형 이중배열 만들기

/*
문제 세분화 해보기..

N이 주어진다.
예시 입출력 배열 int[N][N]이 있다고 할때
배열을 순회하면서 각 요소 추가 해라.. 다만 조건은 시계방향 나선형으로 순회한다.
ex N=4
 1  2  3  4
12 13 14  5
11 16 15  6
10  9  8  7
시작지점은 [0+1][0+1] 각 행과 열의 레이어 한개씩 증가
끝지점은 [0+1][0]
*/

var memo [4][4]int

func gallaxy(n int, memo [4][4]int) [4][4]int {

	memo[0][0] = 1
	memo[0][1] = 2

	//배열의 끝에 다다르거나, 메모이제이션 값이 있을때

	for i := 1; i < n; i++ {
		if i >= 0 && i < n {
			memo[0][i] = memo[0][i-1] + 1
		}
	} //배열의 끝에 다다르거나, 메모이제이션 값이 있을때

	for i := 1; i < n; i++ {
		if i >= 0 && i < n {
			memo[i][n-1] = memo[i-1][n-1] + 1
		}
	} //배열의 끝에 다다르거나, 메모이제이션 값이 있을때

	for i := n - 2; i >= 0; i-- {
		if i >= 0 && i < n {
			memo[n-1][i] = memo[n-1][i+1] + 1
		}
	} //배열의 끝에 다다르거나 메모이제이션 값이 있을때
	for i := n - 2; i > 0; i-- {
		if i >= 0 && i < n {
			memo[i][0] = memo[i+1][0] + 1
		}
	}
	return memo
}
