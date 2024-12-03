package util

import ("bufio"
	"os"
	"time")

// Records the current time and returns a function that will compute and return
// the time elapsed when invoked, it will return a Duration.
func Timer() func() (time.Duration) {
	start := time.Now()
	return func() (time.Duration) {
		return time.Since(start)
	}
}

// Accepts two integers and returns the absolute difference between the two.
func AbsDiff(x int, y int) (int) {
	if (x < y) {
		return y - x
	} else {
		return x - y
	}
}

// Accepts a string with a path to a file and a function that will accept a
// buffered scanner over that file. This function handles the cleanup of the
// opened file.
func LoadTextFile(filename string, scanner_fn func(*bufio.Scanner)) {
	file, error := os.Open(filename)
	if error != nil {
		panic(error)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner_fn(scanner)

	if error := scanner.Err(); error != nil {
		panic(error)
	}
}
