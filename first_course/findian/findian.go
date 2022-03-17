package main

import (
	"fmt"
	"strings"
)

func main() {

	welcomeMessage := `
	Hey! This program will try to find Ian in your input! Amazing!
	`
	fmt.Println(welcomeMessage)
	var someString string
	fmt.Println("Give me some string:")
	fmt.Scan(&someString)

	someStringLowered := strings.ToLower(someString)
	someStringLength := len(someStringLowered)
	firstChar, lastChar := string(someStringLowered[0:1]), string(someStringLowered[someStringLength-1:someStringLength])
	if firstChar == "i" && lastChar == "n" && strings.Contains(someStringLowered, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
