package cmd

import (
	"fmt"
	"log"
	"path/filepath"
	"unicode"

	"github.com/spf13/cobra"
	"github.com/thekad/aoc/2022/pkg/file"
)

func contains(needle rune, haystack string) bool {
	for _, i := range haystack {
		if needle == i {
			return true
		}
	}
	return false
}

func runeValue(r rune) int {
	if unicode.IsLower(r) {
		return int(r) - 96
	}
	return int(r) - 64 + 26
}

var threeCmd = &cobra.Command{
	Use:   "three",
	Short: "Run the third day's exercises",
	Run: func(cmd *cobra.Command, args []string) {
		var filePath string
		var err error

		if len(args) == 0 {
			filePath, err = file.DataFilePath("day-3.txt")
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

		// part 1
		var prioritySum int
		for _, line := range lines {
			middle := len(line) / 2

			left := line[0:middle]
			right := line[middle:]

			for _, needle := range left {
				if contains(needle, right) {
					priority := runeValue(needle)
					prioritySum += priority
					fmt.Println(fmt.Sprintf("Priority for %s is %d (%s)", line, priority, string(needle)))
					break
				}
			}
		}
		fmt.Println("Priority sum is", prioritySum)

		// part 2
		var badgeSum int
		for i := 0; i < len(lines); i += 3 {
			block := lines[i : i+3]

			var badge rune
			for _, needle := range block[0] {
				if !contains(needle, block[1]) {
					continue
				}
				if contains(needle, block[2]) {
					badge = needle
					break
				}
			}
			badgeSum += int(runeValue(badge))
			fmt.Println(fmt.Sprintf("Badge for %s is %s", block, string(badge)))
		}
		fmt.Println("Badge sum is:", badgeSum)
	},
}

func init() {
	rootCmd.AddCommand(threeCmd)
}
