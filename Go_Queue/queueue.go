package main

import (
	"container/list"
	"fmt"
)

// Push value > parameter - *Pointer of List, input any(interface{})
// Push the value at back of list.
func Push(l *list.List, val any) {
	l.PushBack(val) // return the back of list q,
	fmt.Print(l.Back().Value)
}

// Pop the value at the fist of List, and return the element of Value.
func pop(l *list.List) any {
	front := l.Front()
	if front != nil {
		//앞에값 토해냄.
		fmt.Print(front.Value)
		return l.Remove(front)
	}
	return nil
}

func main() {

	var test1 list.List

	//push back 을 사용하여 0부터 입력하여 거꾸로 넣어준당...
	for i := 0; i < 5; i++ {
		Push(&test1, i)

	}
	/*
		//List의 순회,
		for e := test1.Front(); e != nil; e = e.Next() {
			//e의 값을 가져옴
			fmt.Print(e.Value, " ")
		}
	*/
	//먼저 넣은 값부터 튀어나와야함.
	for i := 0; i < 5; i++ {
		pop(&test1)

	}
}
