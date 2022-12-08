package cmd

import (
	"fmt"
	"log"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/thekad/aoc/2022/pkg/file"
	"github.com/thekad/aoc/2022/pkg/stack"
)

func loadStackMap(lines []string) ([]int, map[int]*stack.RuneStack) {
	rss := map[int]*stack.RuneStack{}

	// reverse the array
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}

	// get the headers
	headers := strings.Split(strings.TrimSpace(strings.Replace(lines[0], "   ", " ", -1)), " ")

	// figure out the coordinates starting with the 2nd position, 1 is always 1
	coords := map[int]int{1: 1}
	// we return an ordered key list for convenience
	keys := []int{1}
	rss[1] = stack.NewRuneStack()
	var offset = 3
	for _, h := range headers[1:] {
		header, _ := strconv.Atoi(h)
		keys = append(keys, header)
		rss[header] = stack.NewRuneStack()
		coords[header] = header + offset
		offset += 3
	}

	// skip the header line
	for _, line := range lines[1:] {
		for header, coord := range coords {
			r := rune(line[coord])
			if r != ' ' {
				rss[header].Push(rune(line[coord]))
			}
		}
	}

	sort.Sort(sort.IntSlice(keys))
	return keys, rss
}

var fiveCmd = &cobra.Command{
	Use:   "five",
	Short: "Run the fifth day's exercises",
	Run: func(cmd *cobra.Command, args []string) {
		var filePath string
		var err error

		if len(args) == 0 {
			filePath, err = file.DataFilePath("day-5.txt")
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

		// First part
		head := append([]string{}, chunks[0]...)
		keys, crateMap := loadStackMap(head)
		fmt.Println(fmt.Sprintf("Before: %v", crateMap))

		for _, line := range chunks[1] {
			instruction := strings.Split(line, " ")
			if len(instruction) < 6 {
				break
			}
			qty, _ := strconv.Atoi(instruction[1])
			from, _ := strconv.Atoi(instruction[3])
			to, _ := strconv.Atoi(instruction[5])

			for i := 1; i <= qty; i++ {
				t := crateMap[from].Pop()
				crateMap[to].Push(*t)
			}
		}

		fmt.Println(fmt.Sprintf("After: %v", crateMap))
		fmt.Print("Tip of crate stacks: ")
		for _, key := range keys {
			t := crateMap[key].Pop()
			fmt.Print(fmt.Sprintf("%s", string(*t)))
			crateMap[key].Push(*t)
		}
		fmt.Println()

		// Second part
		head2 := append([]string{}, chunks[0]...)
		keys, crateMap2 := loadStackMap(head2)
		fmt.Println(fmt.Sprintf("Before: %v", crateMap2))

		for _, line := range chunks[1] {
			instruction := strings.Split(line, " ")
			if len(instruction) < 6 {
				break
			}
			qty, _ := strconv.Atoi(instruction[1])
			from, _ := strconv.Atoi(instruction[3])
			to, _ := strconv.Atoi(instruction[5])

			t, err := crateMap2[from].PopMulti(qty)
			if err != nil {
				log.Fatal(err)
			}
			crateMap2[to].Extend(t)
		}

		fmt.Println(fmt.Sprintf("After: %v", crateMap2))
		fmt.Print("Tip of crate stacks: ")
		for _, key := range keys {
			t := crateMap2[key].Pop()
			fmt.Print(fmt.Sprintf("%s", string(*t)))
			crateMap2[key].Push(*t)
		}
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(fiveCmd)
}
