package main

import (
	"fmt"
	"unsafe"
)

type user_int8 struct {
	d int
	b int
	c int
}

// 메모리 패딩을 줄이려면 작은것 부터!!
type user struct {
	user_int8
	a int8 // actual size = 1byte = 8 bits
	e int8
}

// make the codes' dependencies lower and 응집도는 높게,
func main() {
	var k1 user_int8
	var k user = user{
		k1,
		4,
		5,
	}

	fmt.Println(unsafe.Sizeof(k))
	fmt.Println(k.a)
}
