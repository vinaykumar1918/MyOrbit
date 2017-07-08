package main

import "fmt"

func main() {

	a:=43            //set the value of a to 43
	fmt.Println(a)   //print out the value of a
	fmt.Println(&a)  //print out the memory address of a

	var b *int =&a   //"*int" = a pointer to an int, and assigned a the value of a memory address

	fmt.Println(b)   //print out the value of b

	//the above code makes b pointer to the memory address
	//b is of type "int pointer"
	//*int --the * is part of the type --b is of type *int
}


