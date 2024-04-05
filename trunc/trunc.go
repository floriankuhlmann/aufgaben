package main

import (
	"fmt"
	"math"
)

// 3 points will be given if a program is written.
// 2 points will be given if the program compiles correctly.
func main() {

	// Declaring some variables
	var input1 float64
	var input2 float64
	var output1 int
	var output2 int

	// start the input with a message for the user
	fmt.Print("Please insert a float value:\n")
	_, err := fmt.Scanf("%f", &input1)
	if err != nil {
		fmt.Println(err)
		return
	}

	output1 = int(math.Trunc(input1))
	fmt.Printf("The first truncated integer from %f is %d\n", input1, output1)
	// 3 points will be given for the first test execution, if the program correctly prints the truncated integer when a floating point number is entered.

	// continue the input with a message for the user
	fmt.Print("Please insert another float:\n")
	_, err = fmt.Scanf("%f", &input2)
	if err != nil {
		fmt.Println(err)
		return
	}

	output2 = int(math.Trunc(input2))
	fmt.Printf("The second truncated integer from %f is %d\n", input2, output2)
	// 2 points will be given for the second test execution, if the program correctly prints the truncated integer when another floating point number is entered.
	fmt.Print("Thank you.\n")

}
