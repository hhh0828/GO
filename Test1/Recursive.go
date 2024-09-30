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

	ab := 1 // ab는 1의 메모리 공간을 가리키는 이름
	bc := 2 // bc는 2의 메모리 공간을 가리키는 이름
	fmt.Printf("바뀌기 전값 ab = %d, bc = %d\n", ab, bc)
	Swap(&ab, &bc)
	fmt.Printf("바꾸면? ab = %d, bc = %d\n", ab, bc)
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

//메모리 주소 복사. 나는 바보다. ㅋㅋ 한번씩 더 생각해보자. 함수 내부에서는 값복사가 일어나니까.
//메모리 주소를 반환할수 없으니까. 포인터가 가리키는 내부의 값을 변경해주자... 그러면 Return이 없어도 함수는 명시적 으로 동작한다..
//흠 바보 바보바보. 조금 꼼꼼히 보자. 제발좀.
func Swap(ac, bc *int) {
	fmt.Println(ac)
	fmt.Println(bc)
	temp := *ac //여기서는 바뀌는데... ac, bc는 값복사가 발생함.
	*ac = *bc   // 가리키는 공간의 값을 교환 해준다.
	fmt.Println(ac)
	*bc = temp //
	fmt.Println(bc)
}
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

//what i feel currently.

//질문 > 이미지 리사이즈 할때, 서버에서 파일을 받고 식별 시, ... 어떻게 했나 UUID를 저장함.

//계단 오르기 문제.ex N개의 계단을 1개 또는 2개 오른다고 했을때. 그 계단을 오르는 가지수를 구하여라.
//순열과 조합? 알고리즘 공부가 필요함

//시간복잡도 다시 공부 배열 리스트 순회를 할때..그 ..흠.. 문제 이해부터해야함.

//재귀함수 일반화 하는 연습이 필요함. 함수호출 과정에서 스택 즉 쌓여가는 과정에 대해 조금 더 공부가 필요하다. 직관할수있으면 좋은데 어렵네.

//실수했던 리스트. GO에서 지역변수 선언과, 전역변수 선언에 대해....
//함수 선언시 앞글자 대소문자를 확인하는것. 그다음에 필드도 마찬가지임...

//포인터를 다룰때 꼼꼼히 확인해보자... 가장 중요한건 이 포인터가 가르키고있는것 즉 메모리 공간임

//Map은 엄청 많이 다룰것같다. 터커 강의 다시보기. 패키지 뜯어보기.

//손코딩 해보기. // 공책에 직접 코딩을 해보고, 직접 옮겨서 해보기. 알고리즘/자료구조 문제를 풀때 한번 해보자.

//1  두수 스왑 포인터로

//2  재귀함수로 2^N 승수 구하기.

//3 맵을 순회할때 맵의 value값에 접근하여 바꿀때.

//4 전역 변수 지역 변수 외부에서 접근이 가능한가를 코딩할때..

//5 내가썼던 코드들 먼저 둘러봐야할듯..
//6 가비지컬렉터
//힙과 스택영역에서 어떤식으로 메모리회수가 이루어지는가 더 공부해보자......
//7 slice/str 등 레퍼런스타입 한번더 보자.

//10~12 시간복잡도 계산. 배열 순회/ 요소 삭제/추가/ 연속된 계산을 할때 ... 터커 강의 자료구조 다시보기.

//기능 구현과 관련해서 기본적인 조작을 많이 연습해봐야할듯함..

//이제앞을 코딩은 손코딩으로 하고 컴퓨터로 옮기자... 실수를 줄이는데 도움이 될 것 같다. IDE 의존 금지.
//기본기 퍼펙트하게 하자....
