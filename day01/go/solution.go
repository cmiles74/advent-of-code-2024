package main

import ("bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"cmiles74/util")

func load_sample() ([]int, []int){
	return []int{3, 4, 2, 1, 3, 3}, []int{4, 3, 5, 3, 9, 3}
}

func load_input(filename string) ([]int, []int) {
	var list_1 = []int{}
	var list_2 = []int{}

	util.LoadTextFile(filename, func (scanner *bufio.Scanner) {
		for scanner.Scan() {
			values := strings.Fields(scanner.Text())

			value_1, error_1 := strconv.Atoi(values[0])
			value_2, error_2 := strconv.Atoi(values[1])
			if error_1 != nil || error_2 != nil {
				if error_1 != nil {
					panic(error_1)
				} else if error_2 != nil {
					panic(error_2)
				}
			}

			list_1 = append(list_1, value_1)
			list_2 = append(list_2, value_2)
		}
	})

	return list_1, list_2
}

func sort_list(list []int) {
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
}

func count_occurs(sorted_list []int, search_value int) (int) {
	occurs := 0
	for _, value := range sorted_list {
		if value == search_value {
			occurs += 1
		}

		if value > search_value {
			break;
		}
	}
	return occurs
}

func part_1(list_1 []int, list_2 []int) {
	sort_list(list_1)
	sort_list(list_2)

	sum_distance := 0
	for index, value := range list_1 {
		value_1 := value
		value_2 := list_2[index]
		sum_distance += util.AbsDiff(value_1, value_2)
	}

	fmt.Println(sum_distance)
}

func part_2(list_1 []int, list_2 []int) {
	sort_list(list_2)

	sum_occurs := 0
	for _, value := range list_1 {
		sum_occurs += (value * count_occurs(list_2, value))
	}

	fmt.Println(sum_occurs)
}

func main() {
	//list_1, list_2 := load_sample()
	list_1, list_2 := load_input("../input.txt")

	fmt.Println("Part 1 - Sum of Distances:")
	timer_stop := util.Timer()
	part_1(list_1, list_2)
	var part_1_elapsed = timer_stop()

	fmt.Println("\nPart 2 - Sum of Occurences:")
	timer_stop = util.Timer()
	part_2(list_1, list_2)
	var part_2_elapsed = timer_stop()

	fmt.Println("\n----")
	fmt.Printf("Part 1 completed %v\n", part_1_elapsed)
	fmt.Printf("Part 2 completed %v\n", part_2_elapsed)
}
