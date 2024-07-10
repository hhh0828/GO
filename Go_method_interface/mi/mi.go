package main

//interface/fedex를 모듈로 가져옴
import (
	"fmt"
	"interface/fedex"
)

// book 구조체 생성
type book struct {
	name  string
	stock int
}

// 위 북과는 별개
// sender 인터페이스 생성
type Sender interface {
	Send(parcel string)
}

type Postoffice struct {
	///
}

func (po Postoffice) Send(parcel string) {
	fmt.Println(parcel)
}

// send book 함수 생성, 추가로 센더인터페이스를 사용해야하며, 샌더 인터페이스를 가르킬 수 있도록 해당 샌더는 send() 함수를 가지고 있어야 함
func Sendbook(name string, sender Sender) {
	sender.Send(name)
}

func (thebooks book) buycheck(thebook *book, howmanybuy int) {
	thebook.stock = thebook.stock - howmanybuy
}

func main() {
	//sender 선언 fedex의 fedenxsender 구조체 대입
	sender := fedex.Fedexsender{}
	sender2 := Postoffice{}
	var thebook = book{
		name:  "부자되는법",
		stock: 5,
	}
	//senddd 선언 sender 대입

	thebook.buycheck(&thebook, 2)

	fmt.Printf("%s이 보내졌습니다. 남은재고는 %d 권 입니다\n", thebook.name, thebook.stock)

	Sendbook("어린왕자", sender2)
	Sendbook("킹받네!!", sender)
}
