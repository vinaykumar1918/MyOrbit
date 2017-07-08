package main

import "fmt"

func main() {

	a:=43

	fmt.Println(a)    //43
	fmt.Println(&a)   // some memory address of a

	var b *int = &a
	fmt.Println(b)    // print the value of b, which is the memory address of a
	fmt.Println (*b)  // dereferrence, give me the value of the memory address, in this case b is an operator

	//b is an int pointer;
	//b points to the memory address where an int is stored
	//to see the value in that memory address, add a * in front of b
	//this is known as dereferencing
	//the * is an operator in this case
}
