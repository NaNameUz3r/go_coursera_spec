package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (p Animal) Eat() {
	fmt.Println(p.food)
}

func (p Animal) Move() {
	fmt.Println(p.locomotion)
}

func (p Animal) Speak() {
	fmt.Println(p.noise)
}

func isContains(stringSlice []string, checkString string) bool {
	for _, value := range stringSlice {
		if value == checkString {
			return true
		}
	}
	return false
}

func main() {
	animals := make(map[string]Animal, 0)
	animals["cow"] = Animal{"grass", "walk", "moo"}
	animals["bird"] = Animal{"worms", "fly", "peep"}
	animals["snake"] = Animal{"mice", "slither", "hsss"}

	welcomeMessage := `Hello! You can ask what cow, bird or snake eats, how they move and sounds like.
	Enter two words in prompt - name of the animal you interested in (cow, bird, snake),
	and activity (eat, move, speak). For example: cow speak.
	Type exit to stop this madness.

	Have fun and stay safe!
	`
	fmt.Println(welcomeMessage)

	validAnimal := []string{"cow", "bird", "snake"}
	validAction := []string{"Eat", "Move", "Speak"}

	for {
		var animalName string
		var animalAction string
		fmt.Print("> ")
		fmt.Scanf("%s %s", &animalName, &animalAction)

		animalAction = strings.Title(animalAction)

		if animalName == "exit" {
			break
		}

		if !(isContains(validAnimal, animalName) && isContains(validAction, animalAction)) {
			fmt.Println("You messed something, try again")
		}

		switch fetchAction := animalAction; fetchAction {
		case "Eat":
			animals[animalName].Eat()
		case "Move":
			animals[animalName].Move()
		case "Speak":
			animals[animalName].Speak()
		}

	}
}
