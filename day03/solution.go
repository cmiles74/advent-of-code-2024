package main

import ("bufio"
	"fmt"
	"regexp"
	"strconv"
	// "strings"
	"cmiles74/util"
)

func load_sample() ([]string) {
	return []string{
		"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"}
}

func load_input(filename string) ([]string) {
	var segments = []string{}
	util.LoadTextFile(filename, func (scanner *bufio.Scanner) {
		for scanner.Scan() {
			segment := scanner.Text()
			segments = append(segments, segment)
		}
	})

	return segments
}

func scan_memory(memory string) ([]string) {
	instructions := []string{}
	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	for _, instruction := range re.FindAllString(memory, -1) {
		instructions = append(instructions, instruction)
	}

	return instructions
}

func scan_memory_all_instructions(memory string) ([]string) {
	instructions := []string{}
	re := regexp.MustCompile(`(mul\([0-9]{1,3},[0-9]{1,3}\))|(do\(\))|(don't\(\))`)
	for _, instruction := range re.FindAllString(memory, -1) {
		instructions = append(instructions, instruction)
	}

	return instructions
}

func execute_instruction(instruction string) (int) {
	re := regexp.MustCompile(`\d+`)
	values := re.FindAllString(instruction, 2)
	value_left, _ := strconv.Atoi(values[0])
	value_right, _ :=strconv.Atoi(values[1])
	return value_left * value_right
}

func process_memory(debug bool, scan_mem_fn func(string) ([]string),
	memory []string) (int) {

	var output = 0
	var dont_flag = false
	for _, segment := range memory {
		if debug {
			fmt.Println("Memory Segment:\n", segment)
		}

		instructions := scan_mem_fn(segment)
		if debug {
			fmt.Println("Parsed Instructions:\n", instructions)
		}

		if debug {
			fmt.Print("Executed:")
		}
		for _, instruction := range instructions {

			if instruction == "don't()" {
				dont_flag = true
			} else if instruction == "do()" {
				dont_flag = false
			} else {
				if dont_flag == false {
					output += execute_instruction(instruction)

					if debug {
						fmt.Print(" ", execute_instruction(instruction))
					}
				}
			}
		}
		if debug {
			fmt.Println("\n")
		}
	}
	return output
}

func part_1(memory []string) (int) {
	return process_memory(false, scan_memory, memory)
}

func part_2(memory []string) (int) {
	return process_memory(false, scan_memory_all_instructions, memory)
}

func main() {
	//memory := load_sample()
	memory := load_input("input.txt")

	timer_stop := util.Timer()
	var value = part_1(memory)
	var part_1_elapsed = timer_stop()
	fmt.Println("Part 1 - Executed Valid Memory")
	fmt.Println(value)

	timer_stop = util.Timer()
	value = part_2(memory)
	var part_2_elapsed = timer_stop()
	fmt.Println("\nPart 2 - Executed Valid Memory, All Instructions")
	fmt.Println(value)

	fmt.Println("\n----")
	fmt.Printf("Part 1 completed %v\n", part_1_elapsed)
	fmt.Printf("Part 2 completed %v\n", part_2_elapsed)
}
