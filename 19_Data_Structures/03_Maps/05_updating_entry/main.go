package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Tim" : "Good Morning",
		"Jenny": "Bonjour",
	}

	myGreeting["Harleen"] = "howdy"
	fmt.Println(myGreeting)
	myGreeting["Harleen"] = "Gidday"
	fmt.Println(myGreeting)
}
