package fedex

import "fmt"

type Fedexsender struct {
}

func (f Fedexsender) Send(parcel string) {
	fmt.Println("fedex에 의해 배송 완료되었습니다.ㅋ", parcel)
}
