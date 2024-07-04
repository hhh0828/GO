package main

import "fmt"

type Colortype int

const (
	Red Colortype = iota
	Blue
	Green
	Yellow
)

func returnthecolor(color Colortype) (r string) {
	switch color {
	case Red:
		r = "Red"
		fmt.Println(r)
	case Blue:
		r = "Blue"
		fmt.Println(r)
	case Green:
		r = "Green"
		fmt.Println(r)
	}
	return
}

func main() {

	returnthecolor(Green)
}
