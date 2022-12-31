package main

import (
	"fmt"
	"os"
	"strings"
)

// Returns the points from doing a shape
// Used by both ScoreFromOneGame and ScoreFromRoundEnd
func ShapePoints(shape string) int {
	return map[string]int{
		"X": 1, // rock
		"Y": 2, // paper
		"Z": 3, // scissors
	}[shape]
}

// Returns the score from the shape played by the opponent and by us
// Used by p1
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
	move_points := ShapePoints(response)
	game_points := game_win_map[move+" "+response]
	return game_points + move_points
}

// From a move and a game outcome, return the shape that needs to be played
func GetResponseFromGameEnd(move string, game_should_end string) string {
	return map[string]string{
		"A X": "Z", // rock - lose = scissors
		"A Y": "X", // rock - draw = rock
		"A Z": "Y", // rock - win = paper
		"B X": "X", // paper - lose = rock
		"B Y": "Y", // paper - draw = paper
		"B Z": "Z", // paper - win = scissors
		"C X": "Y", // scissors - lose = paper
		"C Y": "Z", // scissors - draw = scissors
		"C Z": "X", // scissors - win = rock
	}[move+" "+game_should_end]
}

// Returns the score from a shape by the opponent and a game outcome
func ScoreFromRoundEnd(move string, game_should_end string) int {
	game_end_map := map[string]int{
		"X": 0, // lose
		"Y": 3, // draw
		"Z": 6, // win
	}
	game_points := game_end_map[game_should_end]
	response_shape := GetResponseFromGameEnd(move, game_should_end)
	move_points := ShapePoints(response_shape)
	return game_points + move_points
}

type part_function func(string, string) int

// Goes through the file and runs the desired function for each line then adds up total points
func GoThroughFile(input string, function part_function) int {
	games := strings.Split(input, "\n")
	total_score := 0
	for _, game := range games {
		moves := strings.Split(game, " ")
		move := moves[0]
		response := moves[1]
		score := function(move, response)
		total_score = total_score + score
		fmt.Println("move", move, "response", response, "score",
			score, "total", total_score)
	}
	return total_score
}

func p1() {
	//file_path := "aoc/d2/test.txt"
	file_path := "aoc/d2/submission.txt"
	file_string, _ := os.ReadFile(file_path)
	total := GoThroughFile(string(file_string), ScoreFromOneGame)
	fmt.Println(total)
}

func p2() {
	//file_path := "aoc/d2/test.txt"
	file_path := "aoc/d2/submission.txt"
	file_string, _ := os.ReadFile(file_path)
	total := GoThroughFile(string(file_string), ScoreFromRoundEnd)
	fmt.Println(total)
}

func main() {
	function_name := os.Args[1]
	fmt.Println(function_name)
	funcs := map[string]func(){
		"p1": p1,
		"p2": p2,
	}
	funcs[function_name]()
}
