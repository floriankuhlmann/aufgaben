package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// You will need to define and use a function called GenDisplaceFn() which takes three float64
// arguments, acceleration a, initial velocity vo, and initial displacement so.
// GenDisplaceFn() should return a function which computes displacement as a function of time,
// assuming the given values acceleration, initial velocity, and initial displacement.

func GenDisplaceFn(a float64, v float64, s float64) func(float64) float64 {

	// The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one
	// float64 argument which is the displacement travelled after time t.
	fn := func(t float64) float64 {
		return 0.5*a*t + v*t + s
	}

	return fn

}

func main() {

	var (
		err          error
		inputmessage = []string{
			"Please insert acceleration:\n",
			"Please insert velocity:\n",
			"Please insert displacement:\n",
			"Please insert time:\n",
		}
		storage = make([]float64, len(inputmessage))
		distfn  func(float64) float64
	)

	// Write a program which first prompts the user
	//to enter values for acceleration, initial velocity, and initial displacement.
	// Then the program should prompt the user to enter a value for time and the
	// program should compute the displacement after the entered time.
	for i := 0; i < len(inputmessage); i++ {

		scanner := bufio.NewScanner(os.Stdin)
		// start the input with a message for the user

		fmt.Print(inputmessage[i])
		scanner.Scan() // use `for scanner.Scan()` to keep reading and then scan for the float value
		storage[i], err = strconv.ParseFloat(scanner.Text(), 64)

		// write errormessage and continue loop if no integer value given
		if err != nil {
			fmt.Print("Input not allowed, please insert float value\n")
			err = nil
			i--
			continue
		}

	}

	// prepare the function a new environment
	distfn = GenDisplaceFn(storage[0], storage[1], storage[2])

	// compute the value and print out
	fmt.Printf("the displacement after %f seconds is: %f \n", storage[3], distfn(storage[3]))

}
