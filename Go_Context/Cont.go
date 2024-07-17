package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	//context의 백그라운드(기본형?) 을 가지고와서, 덮어씌워 새로운 컨텍스트를 만들고 cancel 기능을 반환한다.
	ctx, cancel := context.WithCancel(context.Background())
	// time dead line 설정.  3초뒤에 ctx.done() 발생
	//ctx1, cancel2 := context.WithTimeout(context.Background(),3*time.Second)
	// withvalue 넘버라는 키와 9라는 값
	//ctx2 := context.WithValue(context.Background(), "number", 9)
	go PrintEverySecond(ctx)
	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()
}

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {

		select {
		//context가 끝나면 Done을 반납한다.
		case <-ctx.Done():
			fmt.Println("work done")
			wg.Done()
			return
		case <-tick:
			fmt.Println("tick")
		}

	}

}
