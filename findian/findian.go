package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 3 points will be given if a program is written.
// 2 points will be given if the program compiles correctly.
func main() {
	// Declaring some variables
	var input string
	var hasPrefix bool = false
	var contains bool = false
	var hasSuffix bool = false

	// start the input with a message for the user
	fmt.Print("Please insert a text:\n")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	input = strings.ToLower(scanner.Text())

	if strings.HasPrefix(input, "i") {
		hasPrefix = true
	} else {
		hasPrefix = false
	}

	if strings.ContainsAny(input , "a") {
		contains = true
	} else {
		contains = false
	}

	if strings.HasSuffix(input , "n") {
		hasSuffix = true
	} else {
		hasSuffix = false
	}

	// 3 points will be given for the first test execution,
	//if the program correctly prints "Found!" when a string
	//that contains the characters ‘i’, ‘a’, and ‘n’, such as "iaaaan" is entered.
	if hasPrefix && contains && hasSuffix {
		fmt.Println( "Found!")
		return
	}

	// 2 points will be given for the second test execution,
	//if the program correctly prints "Not Found!" when a
	//string that does not contain the characters ‘i’, ‘a’, and ‘n’ is entered.
	fmt.Println( "Not Found!")
}
