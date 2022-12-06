package cmd

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/spf13/cobra"
)

func allDifferent(s string) bool {
	checks := map[rune]bool{}
	for _, r := range s {
		if checks[r] {
			return false
		}
		checks[r] = true
	}
	return true
}

func findMarker(stream string, packetSize int) (pos int, packet string) {
	for k := range stream {
		if k+packetSize > len(stream) {
			break
		}
		piece := stream[k : k+packetSize]
		if allDifferent(piece) {
			return k + packetSize, piece
		}
	}

	return 0, ""
}

var sixCmd = &cobra.Command{
	Use:   "six",
	Short: "Run the sixth day's exercises",
	Run: func(cmd *cobra.Command, args []string) {
		var filePath string
		var err error

		if len(args) == 0 {
			filePath, err = dataPath("day-6.txt")
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

		lines, err := readLines(filePath)

		if err != nil {
			log.Fatal(err)
		}

		stream := lines[0]

		// part 1
		pos, piece := findMarker(stream, 4)
		fmt.Println(fmt.Sprintf("At character %d found start-of-packet marker: %s", pos, piece))

		// part 2
		pos, piece = findMarker(stream, 14)
		fmt.Println(fmt.Sprintf("At character %d found start-of-message marker: %s", pos, piece))
	},
}

func init() {
	rootCmd.AddCommand(sixCmd)
}
