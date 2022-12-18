package main

import (
	"fmt"
	"os"
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

func main() {
	// file_path := "aoc/d1/test1.txt"
	file_path := "aoc/d1/submission.txt"
	file, _ := os.ReadFile(file_path)
	// fmt.Println(string(file))
	max_cal := find_max_cal(string(file))
	fmt.Println(max_cal)
}
