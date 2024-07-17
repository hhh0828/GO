package main

import (
	"fmt"
	"sync"
	"time"
)

//채널은 고루틴끼리 메시지를 전달할수 있는 메시지 큐 // Thread Safe Que . in the multiThread, it doesn't require lock() feature.
//채널의 크기....
// make(chan int, 2)
// channel에 아무것도 없으면, 데드락 발생함...<< wg.done이 발생하지않고 멈춰있기 때문에, channel 로부터 데이터를 받는 고루틴이 없으면, 받을때까지 기다린다...
// 무한 루프인 경우에 channel을 닫아주면, wg.done으로 빠질수있게 구성할수있음.

func main() {

	var wg sync.WaitGroup

	//채널은 팝 큐,
	ch := make(chan int)
	wg.Add(1)

	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	//ch <- 9
	wg.Wait()

}

// ch는 레퍼런스, 주소값을 받을 필요없다
func square(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second)
	terminate := time.After(10 * time.Second)

	for {
		select {
		//시간 값만 채널에 반환 후, 해당 시간이 반환 되었을때, 함수를 실행한다
		case <-tick:
			fmt.Println("Tick")
		//time.After함수는 채널에 시간값을 반환하고 해당 시간이 지나면 값을 넘겨준다.
		case <-terminate:
			fmt.Println("Termi")
			wg.Done()
			return
			/*
				case n := <-ch:
					fmt.Println("Square:", n*n)
					time.Sleep(time.Second)
			*/
		}

	}
}
