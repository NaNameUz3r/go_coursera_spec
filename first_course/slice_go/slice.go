package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	helloMessage := `
	Hello. The program will constantly ask you to enter an integer on the command line, 
	then add it to the array and print it back as a sorted array. 
	To break the input loop, enter "X" instead of an integer. `

	fmt.Println(helloMessage)
	integerSlice := make([]int, 0)

	for {
		fmt.Println("Please enter integer for sorted array:")
		var userInput string
		fmt.Scan(&userInput)
		if userInput == "X" {
			fmt.Println("Goodbye")
			break
		}
		if userInput, err := strconv.ParseInt(userInput, 10, 64); err == nil {
			integerSlice = append(integerSlice, int(userInput))
			sort.Ints(integerSlice)
			fmt.Println("Current sorted array is:")
			fmt.Println(integerSlice)
		} else if err != nil {
			fmt.Println("This is not integer, try again.")
		}

	}
}
