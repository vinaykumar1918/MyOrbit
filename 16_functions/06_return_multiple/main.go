package main

import "fmt"

func main() {
	fmt.Println(greet("Jane", "Doe", "danial"))

}

func greet (fname, lname, mname string) (string, string) {
	return fmt. Sprint(fname,  lname), fmt.Sprint(mname,  fname)
}