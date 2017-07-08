package _1

import "fmt"

var x int = 43  //x is defined outside of the block and is available all over the package and so called package scope

func main() {
	fmt.Println(x)
	foo()

	y:=17 // Y over hers is defined block level, and is only available inside the block
	fmt.Println(y)

}

func foo() {
	fmt.Println(x)
}