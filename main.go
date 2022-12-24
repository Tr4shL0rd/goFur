package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// fileExists checks if a file exists at the given filepath.
func fileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}

// readFile reads the contents of a file and returns the contents as a string.
// If the file cannot be opened or read, the function will return an error.
func readFile(filepath string) (string, error) {
	bytes, err := ioutil.ReadFile(filepath)
	return string(bytes), err
}

// repeatString repeats a given string n times and returns the result as a new string.
func repeatString(s string, n int) string {
	return strings.Repeat(s, n)
}

func main() {
	// Slice of command line args
	args := os.Args[1:]
	// Loops over each arg and checks if it's a file
	for _, arg := range args {
		// Checks if the file exists
		if fileExists(arg) {
			// Prints "CONTENTS OF <filename>" and prints "-" as a border to separate filename and file contents
			fmt.Printf("CONTENTS OF %s\n%s\n", arg, repeatString("-", len(arg)+12))
			// Actual file contents
			contents, err := readFile(arg)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(contents)
			}
		}
	}
}
