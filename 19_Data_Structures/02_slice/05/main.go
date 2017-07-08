package main

import "fmt"

func main() {

	mySlice := make([]int, 0, 5)
	//mySlice := []int{1, 3, 5, 7, 9, 11,}

	fmt.Println("___________________")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("___________________")

	for i :=0; i < 8; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len:", len(mySlice), "Capacity:", cap(mySlice), "Value:", mySlice[i])
	}
}


//go creates the array which is double the size of the present array, if the present array is not enough