package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var err error
	var temp int
	storage := make([]int, 0, 3)

	for {
		// start the input with a message for the user
		fmt.Print("Please insert value (or enter 'x' to end program):\n")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan() // use `for scanner.Scan()` to keep reading

		// check for 'X' input always first
		input := scanner.Text()
		if input == "X" {
			break
		}

		// scan for the integer value
		temp, err = strconv.Atoi(scanner.Text())

		// write errormessage and continue loop if no integer value given
		if err != nil {
			fmt.Print("Input not allowed, please insert integer value, or stop programm with 'x'\n")
			err = nil
			continue
		}

		// save the input und sort ist
		storage = append(storage, temp)
		sort.Ints(storage)

		// show the input
		fmt.Println("Your input:", storage)
	}
}
