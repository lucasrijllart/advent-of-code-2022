package main

import (
	"fmt"
	"os"
	"strings"
)

// split a string in two using the middle point
func SplitString(input string) []string {
	half := len(input) / 2
	return []string{input[:half], input[half:]}
}

// find the character that exists in two strings
func FindCommonItem(input1 string, input2 string) string {
	var result string
	for _, item := range input1 { // generated by ChatGPT
		if i := strings.Index(input2, string(item)); i != -1 {
			result = string(item)
		}
	}
	return result
}

// return the item priority
// given the value scheme for letters, I think this is a good way of getting the value
// rather than storing every letter in a struct
func GetItemPriority(item string) int {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	value := strings.Index(alphabet, item)
	return value + 1 // offset as a=1 and not 0
}

// calculate total duplicate item value in each rucksack
func GetTotalItemValue(rucksacks []string) int {
	total := 0
	for _, rucksack := range rucksacks {
		compartments := SplitString(rucksack)
		common_item := FindCommonItem(compartments[0], compartments[1])
		total = total + GetItemPriority(common_item)
	}
	return total
}

// calculate total duplicate item value for given file
func p1() {
	//file_path := "aoc/d3/test.txt"
	file_path := "aoc/d3/submission.txt"
	file_string, _ := os.ReadFile(file_path)
	rucksacks := strings.Split(string(file_string), "\n")
	total := GetTotalItemValue(rucksacks)
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
