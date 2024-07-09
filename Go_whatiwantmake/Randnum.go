package main

import (
	"fmt"
	"time"
)

func main() {

	/*
		t := time.Now()
		rand.Seed(t.UnixMicro())
		var randnumb int8 = int8(rand.Intn(100))
		fmt.Println(randnumb)
		var a int8

		var switch1 bool = true
		for i := 0; switch1; i++ {

			fmt.Scan(&a)

			if randnumb == a {
				switch1 = false
				fmt.Println("congragturation you succeed to input the right number we selescted randomly.")
			} else if randnumb > a || randnumb < a {
				fmt.Println("틀렸습니다.")
			}
		}
		t2 := time.Now()
		fmt.Printf("duration time is %v\n", (t2).Sub(t))
		//https://cs.opensource.google/go/go/+/refs/tags/go1.22.5:src/time/time.go;l=135
		//Time Struct와 관련된 문서

		loc, _ := time.LoadLocation("Local")
		const LongForm = "Jan 2, 2006 at 3:04pm"
		fmt.Println(loc)
		t1, _ := time.ParseInLocation(LongForm, "July 9, 2024 at 19:00pm", loc)
		fmt.Println(t1)
	*/
	loc, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		fmt.Println("타임존을 로드하는 데 실패했습니다:", err)
		return
	}
	//const LongForm = "Jan 2, 2006 at 3:04pm"
	//fmt.Println("타임존:", loc)

	const CustomForm = "Jan 2, 2006 at 3:04pm"
	t1, err := time.ParseInLocation(CustomForm, "July 9, 2024 at 7:00pm", loc)
	if err != nil {
		fmt.Println("날짜 문자열을 파싱하는 데 실패했습니다:", err)
		return
	}
	fmt.Println("파싱된 시간:", t1)
}
