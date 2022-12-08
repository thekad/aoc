package cmd

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/thekad/aoc/2022/pkg/file"
)

type move int

const (
	rock     move = 1
	paper    move = 2
	scissors move = 3
)

func (m move) String() string {
	var r string
	switch m {
	case rock:
		r = "rock"
	case paper:
		r = "paper"
	case scissors:
		r = "scissors"
	}

	return r
}

type outcome int

const (
	win  outcome = 6
	draw outcome = 3
	lose outcome = 0
)

func (o outcome) String() string {
	var r string
	switch o {
	case win:
		r = "win"
	case draw:
		r = "draw"
	case lose:
		r = "lose"
	}

	return r
}

func compute(challenge move, response move) outcome {
	switch {
	case ((challenge == rock && response == paper) ||
		(challenge == paper && response == scissors) ||
		(challenge == scissors && response == rock)):
		return win
	case response == challenge:
		return draw
	default:
		return lose
	}
}

func guess(challenge move, needs outcome) move {
	var r move

	switch needs {
	case win:
		switch challenge {
		case rock:
			r = paper
		case paper:
			r = scissors
		case scissors:
			r = rock
		}
	case lose:
		switch challenge {
		case rock:
			r = scissors
		case paper:
			r = rock
		case scissors:
			r = paper
		}
	default:
		r = challenge
	}

	return r
}

var twoCmd = &cobra.Command{
	Use:   "two",
	Short: "Run the second day's exercises",
	Run: func(cmd *cobra.Command, args []string) {
		var filePath string
		var err error
		moves := map[string]move{
			"A": rock,
			"B": paper,
			"C": scissors,
			"X": rock,
			"Y": paper,
			"Z": scissors,
		}
		outcomes := map[string]outcome{
			"X": lose,
			"Y": draw,
			"Z": win,
		}

		if len(args) == 0 {
			filePath, err = file.DataFilePath("day-2.txt")
			if err != nil {
				log.Fatal(err)
			}
		} else {
			filePath, err = filepath.Abs(args[0])
			if err != nil {
				log.Fatal(err)
			}
		}
		log.Println(fmt.Sprintf("Loading file %s", filePath))
		lines, err := file.ReadLines(filePath)
		if err != nil {
			log.Fatal(err)
		}

		var totalScore int

		// first part
		for _, line := range lines {
			challenge := moves[strings.Split(line, " ")[0]]
			response := moves[strings.Split(line, " ")[1]]
			result := compute(challenge, response)

			totalScore += int(response) + int(result)
		}

		fmt.Println("My final score is:", totalScore)

		// second part
		totalScore = 0
		for _, line := range lines {
			challenge := moves[strings.Split(line, " ")[0]]
			needs := outcomes[strings.Split(line, " ")[1]]
			response := guess(challenge, needs)

			totalScore += int(response) + int(needs)
		}

		fmt.Println("The (right) final score should be:", totalScore)
	},
}

func init() {
	rootCmd.AddCommand(twoCmd)
}
