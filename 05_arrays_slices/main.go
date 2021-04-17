package main

import "fmt"

func main() {
	// Arrays

	/* 	var fruitsArr [2]string
	   	// Assign values

	   	fruitsArr[0] = "Apple"
	   	fruitsArr[1] = "Orange" */

	// Declare and Assign

	/* fruitsArr := [2]string{"Apple", "Orange"} */

	fruitsSlice := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(fruitsSlice)
	fmt.Println(len(fruitsSlice))
	fmt.Println(fruitsSlice[1:3])
}
