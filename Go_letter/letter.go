package main

import "fmt"

func main() {

	name := "hhhhh_ 월드"
	//4byte 배열로 변환
	arr := []rune(name)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%T,%d,%c\n", arr[i], arr[i], arr[i])
		fmt.Println(len(arr))
		fmt.Println(len(name))
	}
	for _, v := range arr {

		fmt.Printf("%T,%d,%c\n", v, v, v)

	}
}
