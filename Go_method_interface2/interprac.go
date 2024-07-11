package main

import "fmt"

type Attacker interface {
	attack()
	finish()
}
type status struct {
	str         int
	inteligence int
	dex         int
	Luk         int
}

type human struct {
	status
	name string
	age  int
}

/*
type animal struct {
	status
	name string
	age  int
}
*/
type troll struct {
	status
	name string
	age  int
	adr  int // address the monster is living
}

func (h human) attack() {
	fmt.Println("attack")
}
func (h human) finish() {
	fmt.Println("finish다!!")
}

func (t troll) attack() {
	fmt.Println("this is troll attack!!")
}

func (t troll) finish() {
	fmt.Println("this is troll finish!!")
}

/*
func (a animal) attack() {
	fmt.Println("this is animal attack")
}
*/
//attacker의 구현 함수, 인터페이스를 인자로 받음, 인터페이스는
func attack(a Attacker) {
	a.attack()
}
func finish(a Attacker) {
	a.finish()
}

func main() {

	var human1 human = human{
		status: status{10, 100, 5, 100},
		name:   "King",
		age:    22,
	}
	/*
		var animal1 animal = animal{
			status: status{10, 100, 5, 100},
			name:   "anm",
			age:    3,
		}
	*/
	var trol1 troll = troll{
		status: status{100, 1, 10, 1},
		name:   "tr",
		age:    30,
		adr:    1,
	}
	//attacker1 의 인터페이스는 8byte 짜리 human의 주소와, human의 타입을 가지고있음
	//attacker1이 human1의 주소를 가지고있기 때문에 attack공격시 human의공격이 발현됨.
	attack(human1)
	attack(trol1)
	finish(trol1)
	fmt.Println("1")
	//var attacker3 Attacker
	//attacker3 = animal1
	//attacker3.attack()
}
