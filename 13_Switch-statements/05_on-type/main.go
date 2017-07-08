package main

import "fmt"

func main() {

	// Switch on types
	// --normally we switch on value of variable
	// -go allows you to switch on type of variable

	type Contact struct {
		greeting string
		name 	 string
	}

	func SwitchOnType (x interface{}) {
		switch x.(type){
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		case Contact:
			fmt.Println("contact")
		default:
			fmt.Println("unknown")



		}
	}
}
