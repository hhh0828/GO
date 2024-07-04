package main123

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a int
	var b int
	var c int
	//var msg string = "hello variable"
	stdin := bufio.NewReader(os.Stdin)

	//a = 20

	n, err := fmt.Scanln(&a, &b, &c)

	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n')
	} else {
		fmt.Println(n, "개 입력받았다")
		fmt.Println(a, b, c)
	}

}
