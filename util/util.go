package util

import ("bufio"
	"os"
	"time")

func Timer() func() (time.Duration) {
	start := time.Now()
	return func() (time.Duration) {
		return time.Since(start)
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
