package cmd

import (
	"fmt"
	"log"
	"math"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/thekad/aoc/2022/pkg/file"
)

type direction string

const (
	up    direction = "U"
	down  direction = "D"
	left  direction = "L"
	right direction = "R"
)

func (d direction) String() string {
	switch d {
	case up:
		return "up"
	case down:
		return "down"
	case left:
		return "left"
	case right:
		return "right"
	}

	return ""
}

type distance struct {
	x int
	y int
}

type position struct {
	x int
	y int
}

type instruction struct {
	dir direction
	qty int
}

func tracePath(from position, i instruction) []position {
	steps := []position{}

	switch i.dir {
	case up:
		for step := from.y + 1; step <= from.y+i.qty; step++ {
			steps = append(steps, position{x: from.x, y: step})
		}
	case down:
		for step := from.y - 1; step >= from.y-i.qty; step-- {
			steps = append(steps, position{x: from.x, y: step})
		}
	case right:
		for step := from.x + 1; step <= from.x+i.qty; step++ {
			steps = append(steps, position{x: step, y: from.y})
		}
	case left:
		for step := from.x - 1; step >= from.x-i.qty; step-- {
			steps = append(steps, position{x: step, y: from.y})
		}
	}

	return steps
}

func (p *position) isBordering(to position) bool {
	return math.Abs(float64(to.x-p.x)) < 2.0 && math.Abs(float64(to.y-p.y)) < 2.0
}

func (p *position) move(x, y int) {
	fmt.Println(fmt.Sprintf("  Moving to %d, %d", x, y))
	p.x = x
	p.y = y
}

func (p *position) follow(to position, dir direction) {

	var moveX, moveY int
	move := true

	switch {
	case p.isBordering(to):
		fmt.Println("  No need to move")
		move = false
	case p.y == to.y: // same row
		moveY = p.y
		switch {
		case p.x > to.x:
			fmt.Println("    to my left...")
			moveX = p.x - 1
		case p.x < to.x:
			fmt.Println("    to my right...")
			moveX = p.x + 1
		}
	case p.x == to.x: // same column
		moveX = p.x
		switch {
		case p.y > to.y:
			fmt.Println("    below me...")
			moveY = p.y - 1
		case p.y < to.y:
			fmt.Println("    above me...")
			moveY = p.y + 1
		}
	case to.x > p.x && to.y > p.y:
		fmt.Println("    on my top right corner...")
		moveX = p.x + 1
		moveY = p.y + 1
	case to.x > p.x && to.y < p.y:
		fmt.Println("    on my bottom right corner...")
		moveX = p.x + 1
		moveY = p.y - 1
	case to.x < p.x && to.y > p.y:
		fmt.Println("    on my top left corner...")
		moveX = p.x - 1
		moveY = p.y + 1
	case to.x < p.x && to.y < p.y:
		fmt.Println("    on my bottom left corner...")
		moveX = p.x - 1
		moveY = p.y - 1
	}
	if move {
		p.move(moveX, moveY)
	}
}

var nineCmd = &cobra.Command{
	Use:   "nine",
	Short: "Run the ninth day's exercises",
	Run: func(cmd *cobra.Command, args []string) {
		var filePath string
		var err error

		if len(args) == 0 {
			filePath, err = file.DataFilePath("day-9.txt")
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
		rope := make([]position, 2) // rope of 2 knots
		head := &rope[0]            // pointer for readability
		tail := &rope[1]            // pointer for readability
		tailCounter := map[position]bool{}
		instructions := []instruction{}

		for _, line := range lines {
			parts := strings.Split(line, " ")
			qty, _ := strconv.Atoi(parts[1])
			dir := direction(parts[0])
			instr := instruction{
				dir: dir,
				qty: qty,
			}
			instructions = append(instructions, instr) // store for pt 2
			steps := tracePath(*head, instr)
			fmt.Println(fmt.Sprintf("Path traced: %v %s", steps, dir))
			for _, step := range steps {
				fmt.Println("Head moves")
				head.move(step.x, step.y)
				fmt.Println("Tail follows")
				tail.follow(*head, instr.dir)
				tailCounter[*tail] = true
			}
		}
		fmt.Println(fmt.Sprintf("The tail landed on %d spaces at least once", len(tailCounter)))
		fmt.Println("************************")

		// second part
		rope = make([]position, 10) // rope of 10 knots
		head = &rope[0]             // pointer for readability
		tail = &rope[9]             // pointer for readability
		tailCounter = map[position]bool{}

		for _, instr := range instructions {
			steps := tracePath(*head, instr)
			fmt.Println(fmt.Sprintf("Path traced: %v %s", steps, instr.dir))
			for _, step := range steps {
				fmt.Println("Head moves")
				head.move(step.x, step.y)
				prevKnot := head
				for idx := 1; idx < len(rope); idx++ {
					knot := &rope[idx]
					fmt.Println(fmt.Sprintf("Knot %v follows", idx))
					knot.follow(*prevKnot, instr.dir)
					if knot == tail {
						tailCounter[*tail] = true
					}
					prevKnot = knot
				}
			}
		}
		fmt.Println(fmt.Sprintf("The tail landed on %d spaces at least once", len(tailCounter)))
	},
}

func init() {
	rootCmd.AddCommand(nineCmd)
}
