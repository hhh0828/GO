package m123

import "fmt"

type user struct {
	name   string
	age    int
	gender bool // true = male, false = female.
}

func nama() {

	var user1 user = user{
		name:   "king",
		age:    32,
		gender: true,
	}

	fmt.Printf("username:%s, userage:%d, usergender:", user1.name, user1.age)
	fmt.Println(user1.gender)
}
