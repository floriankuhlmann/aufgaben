package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// 3 points will be given if a program is written.
func main() {

	// Declaring some variable
	inputData := make(map[string]string)
	keys := []string{"name", "addresse"}
	scanner := bufio.NewScanner(os.Stdin)

	for i := range keys {
		// start the input with a message for the user
		fmt.Printf("Please insert your %s:\n", keys[i])
		scanner.Scan()                      // read
		inputData[keys[i]] = scanner.Text() // process
	}

	jsonObj, _ := json.Marshal(inputData)
	fmt.Print("Output data as json:")
	// 5 points will be given if the program correctly prints a JSON object with keys ("name", "address")
	// and they are associated with the name and address that was entered.
	fmt.Println(string(jsonObj))

}

// 2 points will be given if the program compiles correctly.
