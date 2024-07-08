package main

//어우.. .외부로 공개되는건 func에서 case sensitive -
import (
	"usepackage/custompkg"

	"fmt"

	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo/ch16/expkg"
)

func main() {
	custompkg.Printcustom()
	expkg.PrintSample()

	data := []float64{3, 4, 5, 6, 7, 8, 9, 50, 20, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}
