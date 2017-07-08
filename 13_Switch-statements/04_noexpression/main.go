package main

import "fmt"

func main() {

	myFriendname := "Ma"

	switch{
	case len(myFriendname) ==2:
		fmt.Println("Wassup  my friend with name of length 2")
	case myFriendname == "Tim":
		fmt.Println("Wassup TIm")
	case myFriendname =="jenny":
		fmt.Println("Wassup Jenny")
	case myFriendname =="Marcus", myFriendname == "Medhi":
		fmt.Println("your name is either marcus or medhi")
	case myFriendname=="March":
		fmt.Println("wassup mar")
	}
}
