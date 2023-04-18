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

// read file, execute moves and print the top container of each stack
func p1() {
	//file_path := "aoc/d5/test.txt"
	file_path := "aoc/d5/submission.txt"
	file, _ := os.ReadFile(file_path)
	containers, moves := parseInput(string(file))
	movedContainers := MoveContainers(containers, moves)
	topContainers := GetTopContainers(movedContainers)
	fmt.Println(topContainers)
}

// read file, execute moves with 9001 config and print top container of each stack
func p2() {
	//file_path := "aoc/d5/test.txt"
	file_path := "aoc/d5/submission.txt"
	file, _ := os.ReadFile(file_path)
	containers, moves := parseInput(string(file))
	movedContainers := MoveContainers9001(containers, moves)
	topContainers := GetTopContainers(movedContainers)
	fmt.Println(topContainers)
}

// convert the input file into Go data structures. Example:
// A
// B C
// 1 2
// move 1 from 1 to 2
// ==
// containers == [
//
//	["B", "A"]
//	["C"]
//
// ]
// moves == [ [1,1,2] ]
func parseInput(file_input string) ([][]string, [][]int) {
	parts := strings.Split(file_input, "\n\n")
	containersList, movesList := strings.Split(parts[0], "\n"), strings.Split(parts[1], "\n")

	var containers [][]string
	var moves [][]int

	// looping through every 4th number starting with 1 so 1,5,9,... as that's the string index of where the stack is
	for stackIndex := 1; stackIndex < len(containersList[0]); stackIndex = stackIndex + 4 {
		collectedContainers := []string{}
		// go through stack reversed and skipping the first, so the number is skipped and then we move up
		for containerIndex := len(containersList) - 2; containerIndex >= 0; containerIndex-- {
			// if the value is a container add to the collected ones this run
			if value := string(containersList[containerIndex][stackIndex]); value != " " {
				collectedContainers = append(collectedContainers, value)
			}
		}
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

// given containers and a set of moves, execute the moves given and return the final containers
func MoveContainers(containers [][]string, moves [][]int) [][]string {
	for _, move := range moves {
		containersToMove, from, to := move[0], move[1]-1, move[2]-1 // add offset to from
		for moving := 0; moving < containersToMove; moving++ {
			container := containers[from][len(containers[from])-1]        // get value
			containers[from] = containers[from][:len(containers[from])-1] // remove value
			containers[to] = append(containers[to], container)
		}
	}
	return containers
}

// return the top containers of each stack
func GetTopContainers(containers [][]string) string {
	var topContainers string
	for _, stack := range containers {
		topContainers = topContainers + stack[len(stack)-1]
	}
	return topContainers
}

// given containers and a set of moves, execute the moves given and return the final containers
// same as MoveContainers except this time it moves multiple containers at once rather than one
// at the time.
func MoveContainers9001(containers [][]string, moves [][]int) [][]string {
	for _, move := range moves {
		containersToMove, from, to := move[0], move[1]-1, move[2]-1 // add offset to from
		stack := containers[from]
		currentMove := stack[len(stack)-containersToMove:]
		containers[from] = stack[:len(stack)-containersToMove] // remove value
		containers[to] = append(containers[to], currentMove...)
	}
	return containers
}
