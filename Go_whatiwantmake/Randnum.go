package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var randnumb int8 = int8(rand.Intn(100))
	fmt.Println(randnumb)
	var a int8

	var switch1 bool = true
	for i := 0; switch1; i++ {

		fmt.Scan(&a)

		if randnumb == a {
			switch1 = false
			fmt.Println("congragturation you succeed to input the right number we selescted randomly.")
		} else if randnumb > a || randnumb < a {
			fmt.Println("틀렸습니다.")
		}
	}

}
