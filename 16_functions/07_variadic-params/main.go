package main

import "fmt"

func main() {
	n:= average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)

}

func average(sf ...float64) float64{
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	var total float64
	//total :=0.0                //cummilative values
	for_, v := range sf {
		total += v
	}
	return total/ float64(9len(sf))
}