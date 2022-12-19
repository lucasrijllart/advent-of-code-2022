package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func total_calories_by_elf(input []string) map[int]int {
	elves := make(map[int]int)
	var elf_id int
	elf_id = 1
	total_calories := 0
	for _, calories_str := range input {
		//fmt.Println(calories_str)
		calories, _ := strconv.Atoi(calories_str)
		total_calories = total_calories + calories
		if calories_str == "" {
			elves[elf_id] = total_calories
			elf_id++
			total_calories = 0
		}
	}
	elves[elf_id] = total_calories
	return elves
}

func find_max_cal(input string) int {
	list_of_calories := strings.Split(input, "\n")
	elves := total_calories_by_elf(list_of_calories)

	max_cal := 0
	for _, calories := range elves {
		if calories > max_cal {
			max_cal = calories
		}
	}
	return max_cal
}

func find_top_three(input string) int {
	list_of_calories := strings.Split(input, "\n")
	elves_map := total_calories_by_elf(list_of_calories)
	elves := make([]int, 0, len(elves_map))
	for _, value := range elves_map {
		elves = append(elves, value)
	}

	sort.Ints(elves)
	top_three := elves[len(elves)-3:]
	fmt.Println(top_three)
	total := (top_three[0] + top_three[1] + top_three[2])
	return total
}

func read_file(file_path string) string {
	file, _ := os.ReadFile(file_path)
	return string(file)
}

func d1() {
	// file_path := "aoc/d1/test1.txt"
	file_path := "aoc/d1/submission.txt"
	// fmt.Println(string(file))
	max_cal := find_max_cal(read_file(file_path))
	fmt.Println(max_cal)
}

func d2() {
	file_path := "aoc/d1/submission.txt"
	top_three := find_top_three((read_file(file_path)))
	fmt.Println(top_three)
}

func main() {
	function_name := os.Args[1]
	fmt.Println(function_name)
	funcs := map[string]func(){
		"d1": d1,
		"d2": d2,
	}
	funcs[function_name]()
}
