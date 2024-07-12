package main

import (
	"container/list"
	"fmt"
)

func main() {

	/*
		.. 배열
		연속된 메모리 값을가진 주소를 연속된 메모리형태로 저장
		간결하며 불러오기쉽다..

		..리스트
		메모리상에 var a int, var b int 라고 있다고 치면
		a와 b의 공간이 있으면, a주소 b주소 서로연속되지 않은 형태로 주소를 가지고있음.
		a와 b의 포인터값을 가지고있는 리스트라고 함. 불연속 메모리...
		배열과 함께 가장 기본적인 선형 자료구조 중 하나 .
		하나의 데이터에 다른 하나에 데이터가 연결되어있음(선형구조)
		트리구조 (비선형 데이터) a가 b와c와 연계되어있음

		Big-O 표기법
		알고리즘의 효율성을 나타내는 표기법 중 하나로 가장 많이 쓰인다.
		시간, 공간 상한선
	*/

	//표기법 - N개의 요소가 있는 자료 구조
	//ex) for i=0; i < N; i++  {연산1번} o(N)
	//ex) for * for = o(N^2)
	//O(1)<O(N)<O(N'log_2N)<O(N^2)<O(N^3)
	//nLog2n >> tree구조

	//list의 주소값으로 가지는..
	a := list.New()
	e4 := a.PushBack(4)
	e1 := a.PushFront(1)
	//e4의 주소값을 가지고있음
	a.InsertBefore(3, e4)
	//e1의 주소값을 가지고있음
	a.InsertAfter(2, e1)

	//e는 a의 첫값 , e가 비어있지 않으면, e 다음값을 삽입
	for e := a.Front(); e != nil; e = e.Next() {
		//e의 값을 가져옴
		fmt.Print(e.Value, " ")

	}
	/*list의 형태
	//type Element struct {
		value interface{} >> 데이터를 저장하는 필드 // 모든 타입의 값을 가질 수 있음
		Next *Element // 다음 노드의 주소를 가르킴
		Prev *Element // 이전 노드의 주소를 가르킴
	}
	*/

	//자료구조에서 배열과 리스트의 차이점....

	/*
		List 승리
		맨 앞에 요소 삽입 or 삭제 - 배열 =>>  o(N) for문으로 밀어냄
		맨 앞에 요소 삽입 or 삭제 - 리스트 =>> o(1) create new element and edit the next list *mark,

		Array 승리
		특정 요소 접근 - 배열 =>> o(1) index 이동공식 > 배열시작주소 +(인덱스 * 타입크기) 바로이동 가능!
		특정 요소 접근 - 리스트 =>> o(N) index 이동공식 > Next()로 N(Elements 개수)로 ...

		정리 요소삽입이나 삭제가 많은 경우, 리스트가 유리하고 랜덤액서스일 경우 배열이 유리

		데이터 지역성에대한 설명 ....
		데이터가 인접해 일 수록 캐시 성공류률이 올라가고 성능도 증가한다 (CPU 연산) CPU에서 연산할때 데이터 지역을 가져옴...
		캐시데이터 만큼?
		요소 수가 적을 경우, 리스트보다 배열이 빠름..
	*/
}
