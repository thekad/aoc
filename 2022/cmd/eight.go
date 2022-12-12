package cmd

import (
	"fmt"
	"log"
	"path/filepath"
	"sort"

	"github.com/spf13/cobra"
	"github.com/thekad/aoc/2022/pkg/file"
)

func isVisible(i, j, value int, trees [][]int) bool {
	// all trees around the map are visible
	if i == 0 || j == 0 || i == len(trees)-1 || j == len(trees[0])-1 {
		return true
	}
	var left, right, top, bottom bool

	fmt.Println(fmt.Sprintf("%d (%d, %d) ", value, i, j))

	// from the left?
	for jj := 0; jj < j; jj++ {
		if trees[i][jj] < value {
			left = true
		} else {
			left = false
			break
		}
	}
	fmt.Println(fmt.Sprintf("Visible from the left? %v", left))

	// from the top?
	for ii := 0; ii < i; ii++ {
		if trees[ii][j] < value {
			top = true
		} else {
			top = false
			break
		}
	}
	fmt.Println(fmt.Sprintf("Visible from the top? %v", top))

	// from the right?
	for jj := len(trees[i]) - 1; jj > j; jj-- {
		if trees[i][jj] < value {
			right = true
		} else {
			right = false
			break
		}
	}
	fmt.Println(fmt.Sprintf("Visible from the right? %v", right))

	// from the bottom?
	for ii := len(trees) - 1; ii > i; ii-- {
		if trees[ii][j] < value {
			bottom = true
		} else {
			bottom = false
			break
		}
	}
	fmt.Println(fmt.Sprintf("Visible from the bottom? %v", bottom))

	return left || right || top || bottom
}

func getScenicScore(i, j, value int, trees [][]int) int {
	// all trees around the map end up with a scenic score of 0
	if i == 0 || j == 0 || i == len(trees)-1 || j == len(trees[0])-1 {
		return 0
	}

	var left, top, right, bottom int

	fmt.Println(fmt.Sprintf("%d (%d, %d) ", value, i, j))

	// to the left
	for jj := j - 1; jj >= 0; jj-- {
		left++
		if trees[i][jj] >= value {
			break
		}
	}
	fmt.Println(fmt.Sprintf("Scenic score to the left is %v", left))

	// to the top
	for ii := i - 1; ii >= 0; ii-- {
		top++
		if trees[ii][j] >= value {
			break
		}
	}
	fmt.Println(fmt.Sprintf("Scenic score to the top is %v", top))

	// to the right
	for jj := j + 1; jj < len(trees[i]); jj++ {
		right++
		if trees[i][jj] >= value {
			break
		}
	}
	fmt.Println(fmt.Sprintf("Scenic score to the right is %v", right))

	// to the bottom
	for ii := i + 1; ii < len(trees); ii++ {
		bottom++
		if trees[ii][j] >= value {
			break
		}
	}
	fmt.Println(fmt.Sprintf("Scenic score to the bottom is %v", bottom))

	return left * top * right * bottom
}

var eightCmd = &cobra.Command{
	Use:   "eight",
	Short: "Run the eighth day's exercises",
	Run: func(cmd *cobra.Command, args []string) {
		var filePath string
		var err error

		if len(args) == 0 {
			filePath, err = file.DataFilePath("day-8.txt")
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

		allTrees, err := file.ReadLinesAsInt(filePath)

		if err != nil {
			log.Fatal(err)
		}

		// first part
		visibleTrees := []int{}
		for ii := range allTrees {
			for jj := range allTrees[ii] {
				fmt.Print(fmt.Sprintf("%d ", allTrees[ii][jj]))
			}
			fmt.Println()
		}

		for i, row := range allTrees {
			for j, value := range row {
				if isVisible(i, j, value, allTrees) {
					visibleTrees = append(visibleTrees, allTrees[i][j])
				}
			}
		}

		fmt.Println(fmt.Sprintf("The number of trees visible from at least one direction is %d ", (len(visibleTrees))))

		// second part
		for ii := range allTrees {
			for jj := range allTrees[ii] {
				fmt.Print(fmt.Sprintf("%d ", allTrees[ii][jj]))
			}
			fmt.Println()
		}
		fmt.Println()

		scenicScores := []int{}
		for i, row := range allTrees {
			for j, value := range row {
				scenicScores = append(scenicScores, getScenicScore(i, j, value, allTrees))
			}
		}

		sort.Sort(sort.Reverse(sort.IntSlice(scenicScores)))
		fmt.Println(fmt.Sprintf("The largest available scenic score is %v", scenicScores[0:10]))
	},
}

func init() {
	rootCmd.AddCommand(eightCmd)
}
