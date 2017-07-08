package main

import "fmt"

func main() {
	x:= (greet("Jane", "Doe"))
	fmt.Println(x)
}

func greet (fname string, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	return
}