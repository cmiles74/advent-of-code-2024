package util

import ("fmt"
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
