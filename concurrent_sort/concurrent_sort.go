package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

// Write a program to sort an array of integers.
// The program should partition the array into 4 parts, each of which is sorted by a different goroutine.
// Each partition should be of approximately equal size.
//
// Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.
// The program should prompt the user to input a series of integers.
// Each goroutine which sorts ¼ of the array should print the subarray that it will sort.
// When sorting is complete, the main goroutine should print the entire sorted list.
func main() {

	var wg sync.WaitGroup
	var err error
	var temp int
	storage := make([]int, 0, 3)

	for {
		// start the input with a message for the user
		fmt.Print("Please insert value or enter 's' to start sorting:\n")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan() // use `for scanner.Scan()` to keep reading

		// check for 'X' input always first
		input := scanner.Text()

		if input == "s" {

			if len(storage) >= 4 {
				break
			}

			fmt.Print("Not enough input, please insert at least 4 int values\n")
			continue
		}

		// scan for the integer value
		temp, err = strconv.Atoi(scanner.Text())

		// write errormessage and continue loop if no integer value given
		if err != nil {
			fmt.Print("Input not allowed, please insert integer value\n")
			err = nil
			continue
		}

		// save the input und sort ist
		storage = append(storage, temp)

		// show the
		fmt.Println("Your input:", storage)

	}

	// this should be improved using a for loop but its late, i am very tired and have to work tomorrow morning ...
	// so this should work for now
	part1 := storage[:len(storage)/4]
	part2 := storage[len(storage)/4 : len(storage)/4*2]
	part3 := storage[len(storage)/4*2 : len(storage)/4*3]
	part4 := storage[len(storage)/4*3:]

	// print the between steps
	fmt.Println("input part 1:", part1)
	fmt.Println("input part 2:", part2)
	fmt.Println("input part 3:", part3)
	fmt.Println("input part 4:", part4)

	// now do the bubble work in goroutines
	wg.Add(1)
	go BubbleSort(part1, &wg)
	wg.Add(1)
	go BubbleSort(part2, &wg)
	wg.Add(1)
	go BubbleSort(part3, &wg)
	wg.Add(1)
	go BubbleSort(part4, &wg)

	// wait for all
	wg.Wait()

	// recombine it
	storage = part1
	storage = append(storage, part2...)
	storage = append(storage, part3...)
	storage = append(storage, part4...)

	// and make the big Bubblesort now
	wg.Add(1)
	go BubbleSort(storage, &wg)
	wg.Wait()

	// finaly show the bubblesorted input
	fmt.Println("\nAfter Bubble Sorting")
	for _, val := range storage {
		fmt.Println(val)
	}

}

// BubbleSort
// The BubbleSort() function modifies the slice so that the elements are in sorted order.
func BubbleSort(storage []int, wg *sync.WaitGroup) {
	defer wg.Done()
	len := len(storage)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if storage[j] > storage[j+1] {
				Swap(storage, j)
			}
		}
	}
	fmt.Println("Bubblesort:", storage)
	return
}

// Swap
//The Swap() function should take two arguments, a slice of integers and an index value i which
//indicates a position in the slice. The Swap() function should return nothing, but it should swap
//the contents of the slice in position i with the contents in position i+1.
func Swap(storage []int, j int) {
	storage[j], storage[j+1] = storage[j+1], storage[j]
}
