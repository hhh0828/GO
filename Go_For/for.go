package main

import "fmt"

func main() {
	// 총 5줄을 프린트 할건데, 할때마다 1개씩 증가하는거로, 1번째줄은1개 2번째 줄은2개 3번째줄은 3개 4번째줄은 4개 5번째줄은 5개로 할거야.
	// 줄찍기 1개 별찍기 1개 줄찍는것을 기준으로 별찍기.
	//for i := 0; i < 5; i++ {
	//	for j := 0; j < i+1; j++ {
	//		fmt.Print("*")
	//	}
	//	fmt.Println()
	//}

	//for i := 0; i < 5; i++ {
	//for j := 5; j >= i+1; j-- {
	//	fmt.Printf("*")
	//	}
	//	fmt.Println()

	//	}

	//삼각형의 높이 vertic
	//삼각형의 밑변의 길이 2veritc-1개
	//삼각형 라인 하나당 중심 값 = 2vertic / 2 + 1
	var vertic int = 4
	for i := 0; i < vertic; i++ {
		//점을 찍을 때 중심값 계산
		//fmt.Println(i)
		for j := 1; j <= 2*vertic-1; j++ {

			if j > 2*vertic/2+i || j < 2*vertic/2-i {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}

			//if j == 2*vertic/2+1 {
			//	fmt.Print("*")

			//} else {
			//	fmt.Print(" ")
			//}
		}
		fmt.Println()

	}

}
