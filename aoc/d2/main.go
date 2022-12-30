package main

import (
	"fmt"
	"os"
	"strings"
)

func ScoreFromOneGame(move string, response string) int {
	game_win_map := map[string]int{
		"A X": 3, // rock - rock = draw
		"A Y": 6, // rock - paper = win
		"A Z": 0, // rock - scissors = lose
		"B X": 0, // paper - rock = lose
		"B Y": 3, // paper - paper = draw
		"B Z": 6, // paper - scissors = win
		"C X": 6, // scissors - rock = win
		"C Y": 0, // scissors - paper = lose
		"C Z": 3, // scissors - scissors = draw
	}
	move_points_map := map[string]int{
		"X": 1, // rock
		"Y": 2, // paper
		"Z": 3, // scissors
	}

	game_points := game_win_map[move+" "+response]
	move_points := move_points_map[response]
	return game_points + move_points
}

func RPSWin(input string) int {
	games := strings.Split(input, "\n")
	total_score := 0
	for _, game := range games {
		moves := strings.Split(game, " ")
		move := moves[0]
		response := moves[1]
		score := ScoreFromOneGame(move, response)
		total_score = total_score + score
		//fmt.Println("move", move, "response", response, "score", score, "total", total_score)
	}
	return total_score
}

func p1() {
	//file_path := "aoc/d2/test.txt"
	file_path := "aoc/d2/submission.txt"
	file_string, _ := os.ReadFile(file_path)
	total := RPSWin(string(file_string))
	fmt.Println(total)
}

func main() {
	function_name := os.Args[1]
	fmt.Println(function_name)
	funcs := map[string]func(){
		"p1": p1,
	}
	funcs[function_name]()
}
