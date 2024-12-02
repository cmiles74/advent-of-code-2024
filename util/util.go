package util

import ("fmt"
	"time")

func Timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s completed %v\n", name, time.Since(start))
	}
}
