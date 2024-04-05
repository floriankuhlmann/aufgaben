package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var AnimalData = make(map[string]Animal)

	for {
		// start the input with a message for the user
		fmt.Print("Please insert 'newanimal|query' {name of animal} 'cow|bird|snake:\n")
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan() // use `for scanner.Scan()` to keep reading

		// get the input data
		var temp string
		temp = scanner.Text()

		var input []string
		input = strings.Split(temp, " ")

		// process the input
		pocessAnimalData(input, AnimalData)

	}
}

func pocessAnimalData(input []string, AnimalData map[string]Animal) {

	if len(input) != 3 {
		fmt.Printf("Wrong input data. Please try again.\n")
		return
	}

	if input[0] != "query" && input[0] != "newanimal" {
		fmt.Printf("wrong command \n")
		return
	}

	var cow = Cow{
		food: "grass", locomotion: "walk", noise: "moo",
	}
	var bird = Bird{
		food: "worms", locomotion: "fly", noise: "peep",
	}
	var snake = Snake{
		food: "mice", locomotion: "slither", noise: "hsss",
	}

	if input[0] == "newanimal" {
		switch input[2] {
		case "cow":
			AnimalData[input[1]] = cow
			fmt.Printf("animal created %s\n", input[1])
			break
		case "bird":
			AnimalData[input[1]] = bird
			fmt.Printf("animal created %s\n", input[1])
			break
		case "snake":
			AnimalData[input[1]] = snake
			fmt.Printf("animal created %s\n", input[1])
			break
		default:
			fmt.Printf("no animal created. '%s'-type does not exist\n", input[1])
		}
		return
	}

	if input[0] == "query" {
		for index, Animal := range AnimalData {
			if index == input[1] {
				switch input[2] {
				case "move":
					Animal.Move()
					break
				case "eat":
					Animal.Eat()
					break
				case "speak":
					Animal.Speak()
					break
				default:
					fmt.Printf("%s is not able to %s\n", input[1], input[2])
				}
			}
		}
	}

	return
}

/**
INTERFACE DEFINITION
*/

type Animal interface {
	Eat()
	Move()
	Speak()
}

/**
SATISFYING THE INTERFACE
*/

type Cow struct {
	food       string
	locomotion string
	noise      string
}

func (c Cow) Eat() {
	fmt.Print(c.food + "\n\n")
}

func (c Cow) Move() {
	fmt.Print(c.locomotion + "\n\n")
}

func (c Cow) Speak() {
	fmt.Print(c.noise + "\n\n")
}

type Bird struct {
	food       string
	locomotion string
	noise      string
}

func (b Bird) Eat() {
	fmt.Print(b.food + "\n\n")
}

func (b Bird) Move() {
	fmt.Print(b.locomotion + "\n\n")
}

func (b Bird) Speak() {
	fmt.Print(b.noise + "\n\n")
}

type Snake struct {
	food       string
	locomotion string
	noise      string
}

func (s Snake) Eat() {
	fmt.Print(s.food + "\n\n")
}

func (s Snake) Move() {
	fmt.Print(s.locomotion + "\n\n")
}

func (s Snake) Speak() {
	fmt.Print(s.noise + "\n\n")
}
