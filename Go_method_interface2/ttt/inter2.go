package testetsetest

import "fmt"

//인터페이스의 사이즈는 항상 동일 16byte - 8Byte = 구조체의 주소, 나머지8은 구조체 타입

type Database interface {
	get() // get data
	set() // set data
}

//인터페이스를 구현할 Cdataceter의 구조체
type Cdatacenter struct {
}

type Adatacenter struct {
}

type Wrapper struct {
	cdb Cdatacenter //enbeded feild
}

func (w Wrapper) get() {
	w.cdb.getdata()
}

func (w Wrapper) set() {
	w.cdb.setdata()
}

func (a Adatacenter) get() {
	fmt.Println("interface test 2 get")
}

func (a Adatacenter) set() {
	fmt.Println("interface test 2 set")
}

//C dataceter의 getdata - // currently the method's name is mismatched.
func (cc Cdatacenter) getdata() {
	fmt.Println("bring out the data from  the center")
}
func (cc Cdatacenter) setdata() {
	fmt.Println("bring out the data from  the center")
}

func testetsetest() {

	cd := Wrapper{}
	ad := Adatacenter{}
	var database Database
	database = Wrapper{cd.cdb}

	database.get()

	ad.get()

	cd.get()
	cd.set()

}
