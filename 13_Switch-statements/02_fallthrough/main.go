package main

import "fmt"

func main() {

	switch "Marcus" {
	case "Tim":
		fmt.Println("wassup tim")
		fallthrough
	case "jenny":
		fmt.Println("wassup jenny")
	case "Marcus":
		fmt.Println( "wassup Marcus")
		fallthrough
	case  "Julian":
		fmt.Println("wassup Julian")
		fallthrough
	case "Marcus1":
		fmt.Println("wassup Marcus2")
	}
}
