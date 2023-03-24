package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Used to get item value in GetItemPriority
const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// split a string in two using the middle point
func SplitString(input string) []string {
	half := len(input) / 2
	return []string{input[:half], input[half:]}
}

// find the character that exists in two strings
func FindCommonItem(input1 string, input2 string) string {
	var result string
	for _, item := range input1 {
		if i := strings.Index(input2, string(item)); i != -1 {
			result = string(item)
		}
	}
	return result
}

// return the item priority
func GetItemPriority(item string) int {
	return strings.Index(alphabet, item) + 1 // offset as a=1 and not 0
}

// calculate total duplicate item value in each rucksack
func GetRucksackDuplicateValue(rucksack string) int {
	compartments := SplitString(rucksack)
	common_item := FindCommonItem(compartments[0], compartments[1])
	return GetItemPriority(common_item)
}

// calculate total duplicate item value for given file
func p1() {
	// file_path := "aoc/d3/test.txt"
	file_path := "aoc/d3/submission.txt"
	file, _ := os.Open(file_path)
	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		total = total + GetRucksackDuplicateValue(scanner.Text())
	}
	fmt.Println(total)
}

// not used yet
func p2() {
	//file_path := "aoc/d2/test.txt"
	file_path := "aoc/d2/submission.txt"
	os.ReadFile(file_path)
}

// main function to run the function given as argument
func main() {
	function_name := os.Args[1]
	fmt.Println(function_name)
	funcs := map[string]func(){
		"p1": p1,
		"p2": p2,
	}
	funcs[function_name]()
}
