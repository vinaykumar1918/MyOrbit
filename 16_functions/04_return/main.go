package main

import "fmt"

func main() {
	fmt.Println(greet("Jane", "Doe"))
}

func greet (fname, lname string) string{
	return fmt.Sprint(fname, lname)
}


//return a value from a function
//Sprint formats using the default formats for its operands and returns the resulting string
//spaces are added between operands when neither is a string