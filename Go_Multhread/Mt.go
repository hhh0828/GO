package main

//this is a test code for go routine.
import (
	"fmt"
	"sync"
	"time"
)

// go routine uses OS thread, and it's called light weight thread.
//it's different with thread that other lang provides.
// go routine connect to OS thread and OS thread connect to core and it goes to be binded.

type Account struct {
	Balance int
}

var wg sync.WaitGroup

func SumAtoB(a, b int) {

	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}

	fmt.Printf("%d부터%d까지 더한값은 %d 입니다", a, b, sum)
	wg.Done()
}

/*
func PrintHangul() {

	hangul := []rune{'가', '나', '다', '라', '마'}

	for _, v := range hangul {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c", v)

	}

}

func PrintNumbers() {

	for i := 0; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d", i)

	}

}
*/

var mutex sync.Mutex

//각 루틴이 몇번했는지 개수를 세고,

func DepositeAndWithdraw(account *Account) {

	// 어떤 고루틴에서 작업할때, 다른 고루틴이 접근하지 못하게 막아줌. 주로 공통으로 사용하는. 성능향상 불가...
	// dead lock 문제있음..
	mutex.Lock()
	defer mutex.Unlock()

	if account.Balance < 0 {

		panic(fmt.Sprintf("ballance cannot be lower than 0: %d", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}
func main() {
	var a sync.WaitGroup
	/*
		go PrintHangul()
		go PrintNumbers()
	*/
	account := &Account{10000}

	a.Add(10) //10개의 작업이 끝날때까지 기다림.
	for i := 0; i < 10; i++ {
		go func() {
			for {

				DepositeAndWithdraw(account)
			}

		}()
		a.Done()
		fmt.Printf("current account balance is %d", account.Balance)
	}
	a.Wait()
}
