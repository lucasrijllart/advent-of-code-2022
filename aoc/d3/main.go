package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
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
func FindCommonItem(left, right string) (string, error) {
	for _, rune := range left {
		if strings.ContainsRune(right, rune) {
			return string(rune), nil
		}
	}
	return "", errors.New("unable to find common item")
}

// return the item priority
func GetItemPriority(item string) int {
	return strings.Index(alphabet, item) + 1 // offset as a=1 and not 0
}

// calculate total duplicate item value in each rucksack
func GetRucksackDuplicateValue(rucksack string) int {
	compartments := SplitString(rucksack)
	common_item, _ := FindCommonItem(compartments[0], compartments[1])
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
	functionFlag := flag.String("func", "p1", "The function to call")
	flag.Parse()
	switch *functionFlag {
	case "p1":
		p1()
	case "p2":
		p2()
	default:
		log.Fatal("Provide an existing function dammit")
	}
}
