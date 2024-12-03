package main

import ("bufio"
	"fmt"
	"strconv"
	"strings"
	"cmiles74/util")

func load_sample() ([][]int) {
	return [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9}}
}

func load_input(filename string) ([][]int) {
	var reports = [][]int{}
	util.LoadTextFile(filename, func (scanner *bufio.Scanner) {
		for scanner.Scan() {
			levels := strings.Fields(scanner.Text())
			var report = []int{}
			for _, level := range levels {
				level_int, _ := strconv.Atoi(level)
				report = append(report, level_int)
			}
			reports = append(reports, report)
		}
	})

	return reports
}

func all_increasing(report []int) (bool) {
	var valid = true
	var prior_level = -1
	for _, level := range report {
		if prior_level != -1 {
			if prior_level < level {
				valid = false
				break
			}
		}
		prior_level = level
	}

	return valid
}

func all_decreasing(report []int) (bool) {
	var valid = true
	var prior_level = -1
	for _, level := range report {
		if prior_level != -1 {
			if prior_level > level {
				valid = false
				break
			}
		}
		prior_level = level
	}

	return valid
}

func all_distance(report []int) (bool) {
	var valid = true
	var prior_level = -1
	for _, level := range report {
		if prior_level != -1 {
			var distance = util.AbsDiff(prior_level, level)
			if distance < 1 || distance > 3 {
				valid = false
				break
			}
		}
		prior_level = level
	}

	return valid
}

func report_safe(debug bool, report []int) (bool) {
	var safe = false
	if all_increasing(report) || all_decreasing(report) {
		safe = true
	} else if debug {
		fmt.Println(report, "Unsafe, not monotonic")
	}

	if safe {
		if all_distance(report) {
			safe = true
		} else {
			safe = false
			if debug {
				fmt.Println(report, "Unsafe, distance")
			}
		}
	}

	if safe && debug {
		fmt.Println(report, "Safe")
	}

	return safe
}

func remove_index(report []int, index int) ([]int) {
	report_new := make([]int, 0)
	report_new = append(report_new, report[:index]...)
	report_new = append(report_new, report[index + 1:]...)
	return report_new
}

func report_dampener_safe(debug bool, report []int) (bool) {
	var safe = false;

	if report_safe(false, report) {
		return true
	}

	for index := range len(report) {
		if report_safe(false, remove_index(report, index)) {
			return true
		}
	}

	if debug{
		fmt.Println(report, "Unsafe, even when dropping 1 level")
	}

	return safe
}

func num_dampener_safe_reports(debug bool, reports [][]int) (int) {
	var num_safe = 0
	for _, report := range reports {
		if report_dampener_safe(debug, report) {
			num_safe += 1
		}
	}

	return num_safe
}

func num_safe_reports(debug bool, reports [][]int) (int) {
	var num_safe = 0
	for _, report := range reports {
		if report_safe(debug, report) {
			num_safe += 1
		}
	}

	return num_safe
}


func main() {
	//reports := load_sample()
	reports := load_input("input.txt")

	defer util.Timer("Part 1")()
	var num_safe = num_safe_reports(false, reports)
	fmt.Println("Part 1 - Number of Safe Reports")
	fmt.Println(num_safe)

	defer util.Timer("Part 2")()
	var num_dampener_safe = num_dampener_safe_reports(false, reports)
	fmt.Println("\nPart 2 - Number of Safe Reports with Dampener")
	fmt.Println(num_dampener_safe)

	fmt.Println("\n----")
}
