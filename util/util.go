package util

import ("bufio"
	"fmt"
	"os"
	"time")

func Timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s completed %v\n", name, time.Since(start))
	}
}

func AbsDiff(x int, y int) (int) {
	if (x < y) {
		return y - x
	} else {
		return x - y
	}
}

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
