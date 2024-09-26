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
	arrtest2 := []int{1, 2, 6, 11, 5, 4, 8, 9, 10, 11, 13, 12, 14, 15, 0}

	a := twoSum(arrtest2, 15)
	fmt.Println(a)
	b := TwoMin(arrtest2, 3)
	fmt.Println(b)
}

func twoSum(nums []int, target int) []int {
	// map from int -> int
	m := make(map[int]int)

	// python like enumerate
	for i, num := range nums {

		//여기서 m[15-4]를 넣고 m[11] 값이 있으니까 index 반환해주고 3번째인덱스에있는놈과, 현재인덱스 5를반환해주면값이나옴.
		if index, ok := m[target-num]; ok {
			return []int{index, i}
		}
		//여기서 3번째 인덱스인 11을 넣고 // m[11] = 3 임

		m[num] = i
		fmt.Println(m[i])
	}
	return nil
}

func TwoMin(nums []int, target int) []int {

	//맵은 숫자와.. 인덱스를 담음 숫자가 몇번에 있는지.
	m := make(map[int]int)

	for i, num := range nums {

		if index, ok := m[num-target]; ok {
			return []int{index, i}
		}

		m[num] = i
	}
	return nil
}
