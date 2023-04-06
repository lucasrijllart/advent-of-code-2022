package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// return true if there is a section that is fully contained in another
func CheckOverlap(left, right string) bool {
	leftMin, _ := strconv.Atoi(strings.Split(left, "-")[0])
	leftMax, _ := strconv.Atoi(strings.Split(left, "-")[1])
	rightMin, _ := strconv.Atoi(strings.Split(right, "-")[0])
	rightMax, _ := strconv.Atoi(strings.Split(right, "-")[1])

	if leftMin <= rightMin && leftMax >= rightMax || leftMin >= rightMin && leftMax <= rightMax {
		return true
	}
	return false
}

// return true if there is a section that has any overlap with another section
func CheckAnyOverlap(left, right string) bool {
	leftMin, _ := strconv.Atoi(strings.Split(left, "-")[0])
	leftMax, _ := strconv.Atoi(strings.Split(left, "-")[1])
	rightMin, _ := strconv.Atoi(strings.Split(right, "-")[0])
	rightMax, _ := strconv.Atoi(strings.Split(right, "-")[1])

	pair1 := []int{leftMin, leftMax}
	pair2 := []int{rightMin, rightMax}

	if pair1[0] > pair2[0] {
		pair1, pair2 = pair2, pair1
	}
	return pair1[1] >= pair2[0]
}

func CountFullyContainedSections(assignments []string) int {
	total := 0
	for _, line := range assignments {
		inputs := strings.Split(line, ",")
		contains := CheckOverlap(inputs[0], inputs[1])
		if contains {
			total++
		}
		fmt.Println(line, contains)
	}
	return total
}

func CountAnyOverlap(assignments []string) int {
	total := 0
	for _, line := range assignments {
		inputs := strings.Split(line, ",")
		if CheckAnyOverlap(inputs[0], inputs[1]) {
			total++
		}
	}
	return total
}

// count times assignments fully overlap
func p1() {
	//file_path := "aoc/d4/test.txt"
	file_path := "aoc/d4/submission.txt"
	file, _ := os.Open(file_path)
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	total := CountFullyContainedSections(lines)
	fmt.Println(total)
}

// not used yet
func p2() {
	// file_path := "aoc/d4/test.txt"
	file_path := "aoc/d4/submission.txt"
	file, _ := os.Open(file_path)
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	total := CountAnyOverlap(lines)
	fmt.Println(total)
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
		log.Fatal("Provide an existing function darnit")
	}
}
