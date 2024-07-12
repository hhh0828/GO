package main

//defer는 라스트 IN 퍼스트 OUT 구조를 가진다. 스택 구조임.  후입 선출

import (
	"fmt"
	"os"
)

func main() {
	// 대량작업을 할때, 괜찮을듯!! ex) CSV 파일, 등... 데이터 정리 등 데이터베이스의 불러올때.
	//OS 자원 생성
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("errorm failed to create a txt file.")
	}

	defer fmt.Println("test1 1")
	//OS 자원 종료
	defer f.Close()
	defer fmt.Println("파일을 닫습니다")

	fmt.Println("파일에 Hello World 씁니다.")
	//file 내부에 내용을 씀.
	//Fprint의 경우 io.Writer 인터페이스를 인자로 받음, File 인스턴스는 Struct 이고, Write()함수를 가지고 있기때문에.
	fmt.Fprintln(f, "Hello World")

}
