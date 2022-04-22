package main

import (
	"fmt"
)

type Animal interface {
	Eat()	
	Move() 	
	Speak()
}

type Cow struct {name string}

func (t Cow) Eat() {
	fmt.Println("grass")
}

func (t Cow) Move() {
	fmt.Println("walk")
}

func (t Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {name string}

func (t Bird) Eat() {
	fmt.Println("worms")
}

func (t Bird) Move() {
	fmt.Println("fly")
}

func (t Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {name string}

func (t Snake) Eat() {
	fmt.Println("mice")
}

func (t Snake) Move() {
	fmt.Println("slither")
}

func (t Snake) Speak() {
	fmt.Println("hsss")
}

func isContains(stringSlice []string, checkString string) bool {
	for _, value := range stringSlice {
		if value == checkString {
			return true
		}
	}
	return false
}

var welcome_message = `Hello! You can use three following commands:
1) "newanimal <animal_name> <animal_type:["cow"||"bird"||"snake"]" to create new animal of specific type.
2) "query <animal_name> <information_request:["eat"||"move"||"speak"]" to fetch info about animal by name and information type.
3) "exit" - stop 
You do not need to include animal type or information request in quotes in you promt, just write string with three key words.
`

func main() {

	fmt.Println(welcome_message)

	animal_names := make(map[string]Animal, 0)

	valid_commands := []string{"newanimal", "query", "exit"}
	valid_subcommands := []string{"cow", "bird", "snake", "eat", "move", "speak"}

	for {
		var command string 
		var animal_name string
		var sub_command string
		fmt.Print("> ")
		fmt.Scanf("%s %s %s", &command, &animal_name, &sub_command)

		if command == "exit" {
			fmt.Println("See you! Bye!")
			break
		}

		if !(isContains(valid_commands, command) && isContains(valid_subcommands, sub_command)) {
			fmt.Println("You messed commands somehow, try again")
		}

		if command == "newanimal" {
			animal_type := sub_command
			switch fetch_animal_type := sub_command; fetch_animal_type {
			case "cow":
				new_cow := Cow{animal_name}
				animal_names[animal_name] = new_cow
			case "bird":
				new_bird := Bird{animal_name}
				animal_names[animal_name] = new_bird
			case "snake":
				new_snake := Snake{animal_name}
				animal_names[animal_name] = new_snake
			
			}
			fmt.Printf("Super! Created a new animal named %s with a %s type!\n", animal_name, animal_type)
			continue
		} else {
			switch info_request := sub_command; info_request {
			case "eat":
				animal_names[animal_name].Eat()
			case "move":
				animal_names[animal_name].Move()
			case "speak":
				animal_names[animal_name].Speak()
			}
		}

	}
}