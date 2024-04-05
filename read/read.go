package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type fullName struct {
	// golang has no string with limited size, so lets use a byte array
	fname [20]byte
	lname [20]byte
}

func main() {

	var fileName string
	var fileData = []fullName{}
	var fileHandle *os.File
	var err error
	var fileScanner, inputScanner *bufio.Scanner

	// input filename to open
	inputScanner = bufio.NewScanner(os.Stdin)
	fmt.Print("Please insert file name:\n")
	inputScanner.Scan()            // read
	fileName = inputScanner.Text() // process

	// open file to read
	fileHandle, err = os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	// scan file line by line
	fileScanner = bufio.NewScanner(fileHandle)
	for fileScanner.Scan() {
		// get the line separated by spaces
		lineData := strings.Fields(fileScanner.Text())
		inputLine := fullName{
			readInputMax20Byte(lineData[0]),
			readInputMax20Byte(lineData[1]),
		}
		fileData = append(fileData, inputLine)
	}
	fileHandle.Close()

	for i := range fileData {
		fmt.Printf("Print line no %v\n", i)
		fmt.Printf("fname: %v\n", string(fileData[i].fname[:]))
		fmt.Printf("lname: %v\n", string(fileData[i].lname[:]))
	}

}

func readInputMax20Byte(fileData string) [20]byte {

	// byte array size 20 as buffer
	var fnbuffer [20]byte
	chars := []byte(fileData)
	for i := 0; i < len(chars); i++ {
		if i == 20 {
			break
		}
		fnbuffer[i] = chars[i]
	}
	return fnbuffer
}
