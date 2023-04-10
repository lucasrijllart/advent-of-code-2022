package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
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
	file, _ := ioutil.ReadFile(file_path)
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
	containersStr, _ := parts[0], parts[1]

	var containers [][]string
	var moves [][]int

	fmt.Println(containersStr)

	containersList := strings.Split(containersStr, "\n")
	println(containersList)

	for _, container := range containersList {
		fmt.Println(container)
		for index := 0; index < len(container); index++ {
			if index%2 == 0 {
				fmt.Println(container[index])
			}
		}
	}
	return containers, moves
}
