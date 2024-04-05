package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence of up to 10 integers. The program
// should print the integers out on one line, in sorted order, from least to greatest.
// Use your favorite search tool to find a description of how the bubble sort algorithm works.
func main() {
	var err error
	var storagesize int = 10
	storage := make([]int, storagesize)

	for i := 0; i < storagesize; i++ {
		// start the input with a message for the user
		fmt.Print("Please insert int value:\n")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan() // use `for scanner.Scan()` to keep reading

		// scan for the integer value
		storage[i], err = strconv.Atoi(scanner.Text())

		// write errormessage and continue loop if no integer value given
		if err != nil {
			fmt.Print("Input not allowed, please insert integer value\n")
			err = nil
			continue
		}
	}

	// Bubblesort now
	BubbleSort(storage)

	// show the bubblesorted input
	fmt.Println("\nAfter Bubble Sorting")
	for _, val := range storage {
		fmt.Println(val)
	}

}

// BubbleSort
// write a function called BubbleSort() which takes a slice of integers as an argument and returns nothing.
//The BubbleSort() function should modify the slice so that the elements are in sorted order.
func BubbleSort(storage []int) {
	len := len(storage)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if storage[j] > storage[j+1] {
				Swap(storage, j)
			}
		}
	}
}

// Swap
//The Swap() function should take two arguments, a slice of integers and an index value i which
//indicates a position in the slice. The Swap() function should return nothing, but it should swap
//the contents of the slice in position i with the contents in position i+1.
func Swap(storage []int, j int) {
	storage[j], storage[j+1] = storage[j+1], storage[j]
}
