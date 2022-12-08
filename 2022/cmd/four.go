package cmd

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/thekad/aoc/2022/pkg/file"
)

func rangeToMap(r string) (map[int]bool, error) {
	var a = map[int]bool{}

	start, err := strconv.Atoi(strings.Split(r, "-")[0])

	if err != nil {
		return a, err
	}

	end, err := strconv.Atoi(strings.Split(r, "-")[1])

	if err != nil {
		return a, err
	}

	for i := start; i <= end; i++ {
		a[i] = true
	}

	return a, nil
}

func fullyContains(a, b map[int]bool) bool {
	for i := range b {
		if !a[i] {
			return false
		}
	}

	return true
}

func overlap(a, b map[int]bool) bool {
	for i := range b {
		if a[i] {
			return true
		}
	}

	return false
}

var fourCmd = &cobra.Command{
	Use:   "four",
	Short: "Run the fourth day's exercises",
	Run: func(cmd *cobra.Command, args []string) {
		var filePath string
		var err error

		if len(args) == 0 {
			filePath, err = file.DataFilePath("day-4.txt")
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

		// first part
		var c int
		for _, line := range lines {
			first, err := rangeToMap(strings.Split(line, ",")[0])

			if err != nil {
				log.Fatal(err)
			}

			second, err := rangeToMap(strings.Split(line, ",")[1])

			if err != nil {
				log.Fatal(err)
			}

			// I split these just for printing's sake :P
			if fullyContains(first, second) {
				fmt.Println(fmt.Sprintf("%v fully contains %v", first, second))
				c++
				continue
			}
			if fullyContains(second, first) {
				fmt.Println(fmt.Sprintf("%v fully contains %v", second, first))
				c++
			}
		}

		fmt.Println(fmt.Sprintf("Total of fully contained sets: %d", c))

		c = 0
		for _, line := range lines {
			first, err := rangeToMap(strings.Split(line, ",")[0])

			if err != nil {
				log.Fatal(err)
			}

			second, err := rangeToMap(strings.Split(line, ",")[1])

			if err != nil {
				log.Fatal(err)
			}

			// I split these just for printing's sake :P
			if overlap(first, second) {
				fmt.Println(fmt.Sprintf("%v overlap with %v", first, second))
				c++
				continue
			}
			if overlap(second, first) {
				fmt.Println(fmt.Sprintf("%v overlap with %v", second, first))
				c++
			}
		}

		fmt.Println(fmt.Sprintf("Total of overlapping sets: %d", c))
	},
}

func init() {
	rootCmd.AddCommand(fourCmd)
}
