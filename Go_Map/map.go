package main

import (
	"container/list"
	"fmt"
)

type Product struct {
	Name  string
	Price int
}

const M = 10

type Numbofpr int

func hash(d Numbofpr) Numbofpr {
	return d % M
}

// 상품 코드 입력 시, 반환
func Retrivepr(c Numbofpr, m1 [10]*list.List) Product {

	switch hash(c) {

	case 0:
		for e := m1[0].Front(); e != nil; e = e.Next() {
			//e의 값을 가져옴

		}

		//건강용품
	case 1:
		m1[1].PushFront(b)
		//미국산
	case 2:
		m1[2].PushFront(b)
		//중국산
	case 3:
		m1[3].PushFront(b)
		//식용품
	case 4:
		m1[4].PushFront(b)
		//게임용품
	case 5:
		m1[5].PushFront(b)
		//학생판매용
	case 6:
		m1[6].PushFront(b)
		//특별손님용
	case 7:
		m1[7].PushFront(b)
		//코딩용
	case 8:
		m1[8].PushFront(b)
		//공부용
	case 9:
		m1[9].PushFront(b)
	}

}

func Addpr(m1 [10]*list.List, c Numbofpr, b map[Numbofpr]Product) {

	switch hash(c) {
	//사무용품 상수타입으로 대입해서 사용가능.
	case 0:
		m1[0].PushFront(b)

		//건강용품
	case 1:
		m1[1].PushFront(b)
		//미국산
	case 2:
		m1[2].PushFront(b)
		//중국산
	case 3:
		m1[3].PushFront(b)
		//식용품
	case 4:
		m1[4].PushFront(b)
		//게임용품
	case 5:
		m1[5].PushFront(b)
		//학생판매용
	case 6:
		m1[6].PushFront(b)
		//특별손님용
	case 7:
		m1[7].PushFront(b)
		//코딩용
	case 8:
		m1[8].PushFront(b)
		//공부용
	case 9:
		m1[9].PushFront(b)
	}

}

func main() {

	// map[key]value
	//    Key type / value type
	// ex) map[int]string
	//doesn't guarantee the order of popping value.
	m := make(map[int]Product)
	a := Product{"지우개", 1000}
	m[10] = a

	//m1배열 안에 product타입의 리스트가 elements로 들어가야함..
	m1 := [10]*list.List{}
	for i := range m1 {
		m1[i] = list.New()
	}
	Addpr(m1, Product{"지우개", 1000}, 10)
	Addpr(m1, Product{"호치개스", 2000}, 20)
	Addpr(m1, Product{"붕어빵", 200}, 13)
	Addpr(m1, Product{"잉어빵", 300}, 23)
	Addpr(m1, m, 10)
	//map 순회.

	fmt.Println(m1[0].Front().Value)
	fmt.Println(m1[0].Back().Value)

	for i := 0; i < len(m1); i++ {
		for e := m1[i].Front(); e != nil; e = e.Next() {
			//e의 값을 가져옴
			fmt.Printf("m1의 index : %d 입니다", i)
			fmt.Println(e.Value)

		}
	}

}
