package main

// Oct 2nd, the judgement site has been downed.

//팰린드룸의 일반화
//앞으로 봐도 같고 뒤로 봐도 같은지 확인하기위해서,
//스택이 열리는 시점과 닫히는시점은
//

//일반화해보자...
//a[0] == a[n] if true else false
//a[0+1] == a[n-1]
//a[0+2] == a[n-2] // 중간지점까지는 확인해줘야함.
//a[0+k] == a[n-k] //
//a[n/2] == a[n-n/2] // 여기까지 체크하고 '이상없으면' 리턴 함수가 닫히면서 스택을 전달해줘야하는지.

//함수 호출스택 과정에 고수가되면 재귀함수를 맘대로 쓸수있다. 그림을 그리면서해보자
/*
stack 2 - {"ABA", 1 ,1} return int > 1 아래갚이 스택1로 전달됨.
stack 1 - {"ABA", 0 ,2} return func {} int > 1
*/
//개수를 체크하려면 전역변수를 두고 return 시점전에 한번씩 카운트업 해줄것인가
var count int

func pr(a []byte, i, j int) int {
	if i >= j { // 별일없으면 1전달해주자~
		return 1
	} else if a[i] != a[j] {
		return 0
	}
	//들어가는 시점에서 위의 가정을 제외하고 다른 계산의 가정을 없애야함..
	count++
	return pr(a, i+1, j-1) //함수 스택 추가. //return (tpye int)
	//위 실행 흐름을 보면 입력한 i값이j보다 커질 때 return 1이다...
	//즉 가운데까지 가면서. 서로 체크하면서 같지 않으면 0을 리턴하고 아니면 i가 커져서 1이나올떄까지 계속 검사한다.
}

//편법
//2n+1 = 짝수개의 문자열 스택열리는 횟수
//2n+2 = 홀수개의 문자열 스택열리는 횟수

func prh(a []byte, i, j int) bool {
	if i >= j {
		return true
	}
	if a[i] != a[j] {
		return false
	}
	return prh(a, i+1, j-1)
}
