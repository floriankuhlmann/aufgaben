package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() {
	fmt.Print(a.food + "\n\n")
}

func (a *Animal) Move() {
	fmt.Print(a.locomotion + "\n\n")
}

func (a *Animal) Speak() {
	fmt.Print(a.noise + "\n\n")
}

func main() {

	for {
		// start the input with a message for the user
		fmt.Print("Please insert two values: as first the animal name and as a second value the requested information:\n")
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan() // use `for scanner.Scan()` to keep reading

		var temp string
		temp = scanner.Text()

		var input []string
		input = strings.Split(temp, " ")

		if len(input) == 2 {
			if !checkInput(input) {
				fmt.Printf("no data available for animal '%s' and requested info '%s'\n", input[0], input[1])
			}
		}
	}
}

func checkInput(input []string) bool {

	AnimalData := make(map[string]Animal)
	AnimalData["cow"] = Animal{food: "grass", locomotion: "walk", noise: "moo"}
	AnimalData["bird"] = Animal{food: "worms", locomotion: "fly", noise: "peep"}
	AnimalData["snake"] = Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	for index, Animal := range AnimalData {
		if index == input[0] {
			switch input[1] {
			case "move":
				Animal.Move()
				return true
			case "eat":
				Animal.Eat()
				return true
			case "speak":
				Animal.Speak()
				return true
			}
		}
	}

	return false
}
