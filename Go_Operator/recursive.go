package main

import "fmt"

//this is recursivetest - this scope make stacks and it calls itselfs until this meet a condition of exit code
func Recursivetest(a int) {
	//the condition for the script to escape
	if a == 0 {
		return
	}

	fmt.Println(a)
	Recursivetest(a - 1)
	fmt.Println("this is after returning", a)

}
