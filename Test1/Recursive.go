package main

import "fmt"

func main() {

	a := make(map[string]Patients, 4) // 터커 맵강의 한번 더보자...ㅜ_ㅜ
	//해쉬함수의 배열이고 배열의
	a["1"] = NewPatients("홍", 20, 32)
	a["2"] = NewPatients("김", 20, 33)
	a["3"] = NewPatients("테", 30, 1)
	a["4"] = NewPatients("정", 5, 15)

	mapcycle(&a)
	for _, pat := range a {
		fmt.Println(pat)
	}
}

//실행흐름 - 일반화
//스스로 생각해보자.. 2의 제곱수를 표현하는것  2^n = 2^n-1 * 2 ..였군

func recursive(x, n int) int {
	if n == 0 {
		return 2

	}

	return recursive(x, n-1) * 2
}

//몰랐던 문제 Capture 문제 // >> wg 관련 고루틴 리터럴 함수 관련 외부 변수 캡쳐 다시보기.
//맵을 순회할때 for 의 range문 관련해서...  객체에 직접 접근 ?

func NewPatients(name string, discount, Age int) Patients {
	a := new(Patients)
	a.Age = Age // 포인터지만 접근 가능/편의상
	a.Discount = discount
	a.Name = name
	return *a // 역참조로 받기해보자...
}

type Patients struct {
	Name     string
	Discount int
	Age      int
}

func mapcycle(a *map[string]Patients) map[string]Patients {

	// 맵을 포인터로 받고 맵이 가리키는 공간의 맵을 순회함..
	//Range는 순회시 값복사가 일어남. ...........................................

	//a의 존재하는 갭체를 복사해서 가져옴.
	for k, av := range *a {
		av.Discount = 50 // 복사된 av객체에Discount를 수정해주고
		(*a)[k] = av     // 수정된 값을 다시 맵에 할당.... ㅜㅜ 이런기본적인걸...
	}

	return *a
	/*
		for _, v := range a {
			fmt.Println(v)
		}\
	*/
}

//질문 > 이미지 리사이즈 할때, 서버에서 파일을 받고 식별 시, ... 어떻게 했나 UUID를 저장함.
//
