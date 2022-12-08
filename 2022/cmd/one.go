package cmd

import (
	"fmt"
	"log"
	"path/filepath"
	"sort"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/thekad/aoc/2022/pkg/file"
)

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

var oneCmd = &cobra.Command{
	Use:   "one",
	Short: "Run the first day's exercises",
	Run: func(cmd *cobra.Command, args []string) {
		var filePath string
		var err error

		if len(args) == 0 {
			filePath, err = file.DataFilePath("day-1.txt")
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
		chunks, err := file.ReadChunks(filePath, "\n\n")
		if err != nil {
			log.Fatal(err)
		}

		calories := []int{}

		for _, chunk := range chunks {
			var calorie int
			for _, c := range chunk {
				i, _ := strconv.Atoi(c)
				calorie += i
			}

			calories = append(calories, calorie)
		}

		sort.Sort(sort.Reverse(sort.IntSlice(calories)))

		fmt.Println("Most calories is", calories[0])
		fmt.Println("Top 3 calories are", calories[0:3])
		fmt.Println("Sum of top 3 is", sum(calories[0:3]))
	},
}

func init() {
	rootCmd.AddCommand(oneCmd)
}
