package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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

// not used yet
func p1() {
	file_path := "aoc/d5/test.txt"
	// file_path := "aoc/d4/submission.txt"
	file, _ := os.ReadFile(file_path)
	containers, _ := parseInput(string(file))
	fmt.Println(containers)
}

// not used yet
func p2() {
	// file_path := "aoc/d4/test.txt"
	// file_path := "aoc/d4/submission.txt"
	// file, _ := os.Open(file_path)
	// defer file.Close()
	// var lines []string
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	lines = append(lines, scanner.Text())
	// }
	// total := CountAnyOverlap(lines)
	// fmt.Println(total)
}

func parseInput(file_input string) ([][]string, [][]int) {
	parts := strings.Split(file_input, "\n\n")
	containersList, movesList := strings.Split(parts[0], "\n"), strings.Split(parts[1], "\n")

	var containers [][]string
	var moves [][]int

	//fmt.Println(containersStr)

	// looping through every 4th number starting with 1 so 1,5,9,... as that's the string index of where the stack is
	for stackIndex := 1; stackIndex < len(containersList[0]); stackIndex = stackIndex + 4 {
		//fmt.Println(stackIndex)
		collectedContainers := []string{}
		// go through stack reversed and skipping the first, so the number is skipped and then we move up
		for containerIndex := len(containersList) - 2; containerIndex >= 0; containerIndex-- {
			// if the value is a container add to the collected ones this run
			if value := string(containersList[containerIndex][stackIndex]); value != " " {
				collectedContainers = append(collectedContainers, value)
			}
		}
		//fmt.Printf("CC: %v\n", collectedContainers)
		containers = append(containers, collectedContainers)
	}

	// parse moves
	for _, move := range movesList {
		movesFound := regexp.MustCompile(`\d+`).FindAllString(move, -1)
		lineMove := make([]int, len(movesFound))
		for i, s := range movesFound {
			lineMove[i], _ = strconv.Atoi(s)
		}
		moves = append(moves, lineMove)
	}

	return containers, moves
}
