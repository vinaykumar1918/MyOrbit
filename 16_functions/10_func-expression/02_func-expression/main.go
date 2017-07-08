package main

import "fmt"

func main() {

	greeting := func() {               // when assigning a function to a variable is called func expression
		fmt.Println("Hello World")
	}

	greeting()
}
