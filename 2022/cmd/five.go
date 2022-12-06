package cmd

import (
	"fmt"
	"log"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// RuneStack is a naive stack implementation of type rune
type RuneStack struct {
	elements []rune
}

// NewRuneStack is the constructor
func NewRuneStack() *RuneStack {
	rs := RuneStack{
		elements: []rune{},
	}

	return &rs
}

// Push is a naive push implementation
func (s *RuneStack) Push(e rune) {
	s.elements = append(s.elements, e)
}

// IsEmpty returns true if the stack is empty
func (s *RuneStack) IsEmpty() bool {
	return len(s.elements) == 0
}

// Pop is a naive pop implementation
func (s *RuneStack) Pop() *rune {
	if s.IsEmpty() {
		return nil
	}
	r := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]

	return &r
}

// PopMulti returns a slice of how many items from the top of the stack
func (s *RuneStack) PopMulti(n int) ([]rune, error) {
	r := []rune{}

	if n > len(s.elements) {
		return r, fmt.Errorf("Can't pop %d when len is %d", n, len(s.elements))
	}

	idx := len(s.elements) - n
	r = append(r, s.elements[idx:]...)
	s.elements = s.elements[:idx]

	return r, nil
}

// Extend will append the given array to the existing array
func (s *RuneStack) Extend(add []rune) {
	s.elements = append(s.elements, add...)
}

func (s *RuneStack) String() string {
	var r []string
	for _, e := range s.elements {
		r = append(r, string(e))
	}
	return fmt.Sprintf("%v", r)
}

func loadStackMap(lines []string) ([]int, map[int]*RuneStack) {
	rss := map[int]*RuneStack{}

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
	rss[1] = NewRuneStack()
	var offset = 3
	for _, h := range headers[1:] {
		header, _ := strconv.Atoi(h)
		keys = append(keys, header)
		rss[header] = NewRuneStack()
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
			filePath, err = dataPath("day-5.txt")
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

		chunks, err := readChunks(filePath)

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
