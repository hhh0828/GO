package main

import "fmt"

// 함수 외부에서는 반드시 var을 붙혀야함.
// arr1test := [...]string{"a", "b", "c"}
//선언
//var numbs [5]int = [5]int{1, 2, 3, 4, 5}

func main() {
	arrtest := [...]string{"a", "b", "c"}
	//fmt.Println(len(arrtest))
	for i := 0; i <= len(arrtest)-1; i++ {
		fmt.Println(arrtest[i])

	}

	for _, v := range arrtest {
		println(v)
	}
}
