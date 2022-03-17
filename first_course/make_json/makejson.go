package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string
	Address string
}

func main() {
	welcomeMessage := `
	This program is used for generating json file with is contain pairs of persons names 
	and thier addresses. Enter name and adress in corresponding prompts. You can enter more than
	one Person-Address pair, when you done enter "exit" in promt for printing your json in stdout.`
	fmt.Println(welcomeMessage)

	addressBook := make(map[string]string)

	for {
		personName, personAddress := string, string
		fmt.Println("Enter name, or 'exit' for printing json'ed adrees book:")
		fmt.Scanf(&personName)
		if personName == "exit" {
			jsonedAddressBook := json.Marshal(addressBook)
			fmt.Println(jsonedAddressBook)
		}

	}
}
