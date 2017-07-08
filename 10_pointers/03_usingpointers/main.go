package main

import "fmt"

func main() {

	a:=43

	fmt.Println(a)  //43
	fmt.Println(&a)    //Prints a's memory address

	var b *int = &a
	fmt.Println(b)      // memory address of b
	fmt.Println(*b)     //43

	*b = 42 //b says, "The value at this address, change it to 42
	fmt.Println(a)       //42

	//this is useful, we can pass a memory address instead of a buch of baluesand then we can still change the value of whatever is stored
	//this makes our programs more performant
	//we don't have to pass around large amounts of data
}
