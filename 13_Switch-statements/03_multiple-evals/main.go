package main

import "fmt"

func main() {

	switch "Marcus" {
	case "Tim", "Marcus":
		fmt.Println("wassup tim")
	case "jenny":
		fmt.Println("wassup jenny")
	case "Marcus2":
		fmt.Println( "wassup Marcus")
	case  "Julian":
		fmt.Println("wassup Julian")
	case "Marcus1":
		fmt.Println("wassup Marcus2")
	}
}
