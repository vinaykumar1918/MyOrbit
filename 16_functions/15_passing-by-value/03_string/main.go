package main

import "fmt"

func main() {

	name := "vinay"
	fmt.Println(name)  //vinay

	changeMe(name)

	fmt.Println(name) //vinay
}

func changeMe(z string) {
	fmt.Println(z)   //vinay
	z= "Rocky"
	fmt.Println(z)   //Rocky
}