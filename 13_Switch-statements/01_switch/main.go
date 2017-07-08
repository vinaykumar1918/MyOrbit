package main

import "fmt"

func main() {

	switch "medhi" {
	case "Daniel" :
		fmt.Println("Wassup Daniel")
	case "medhi":
		fmt.Println("Wassup Medhi")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	default:
		fmt.Println("You have no friend")
	}
}
