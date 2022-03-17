package main

import (
	"fmt"
	"strconv"
)

func main() {

	helloMessage := `
	Hello, this is simple program which will take a number 
	from input and print out integer part of this number.
	If you want to exit program type "exit"`
	fmt.Println(helloMessage)

	for {
		var userInput string
		fmt.Println("Please enter the real number:")
		fmt.Scan(&userInput)
		if userInput == "exit" {
			fmt.Println("Have a nice day, goodbye!")
			break
		}
		if floatInput, err := strconv.ParseFloat(userInput, 32); err == nil {
			integerInputPart := int(floatInput)
			fmt.Printf("The integer part of number you typed is %v\n", integerInputPart)
			continue
		} else if err != nil {
			fmt.Printf("Look like wrong input, try again with some number.\n")
		}
	}

}
