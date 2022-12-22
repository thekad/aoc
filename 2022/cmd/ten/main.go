package ten

import (
	"fmt"
	"strconv"
	"strings"
)

type command struct {
	instr  string
	cycles int
	value  int
}

func (c *command) exec(acc int) []int {
	r := []int{}

	for i := 1; i < c.cycles; i++ {
		r = append(r, acc)
	}
	r = append(r, acc+c.value)

	return r
}

type sprite struct {
	pos []int
}

func newSprite() sprite {
	return sprite{
		pos: []int{0, 1, 2},
	}
}

func (s *sprite) covers(x int) bool {
	for _, p := range s.pos {
		if p == x {
			return true
		}
	}

	return false
}

func (s *sprite) move(to int) {
	s.pos = []int{to - 1, to, to + 1}
}

// Main method for cmd "ten"
func Main(lines []string) {
	commands := []command{}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		value := 0
		c := command{}
		if parts[0] == "noop" {
			c.instr = "noop"
			c.cycles = 1
		} else {
			c.instr = "addx"
			c.cycles = 2
			value, _ = strconv.Atoi(parts[1])
		}
		c.value = value
		commands = append(commands, c)
	}

	// first part
	results := []int{}
	reg := 1 // initial value
	for _, cmd := range commands {
		rc := cmd.exec(reg)
		reg = rc[len(rc)-1]
		results = append(results, rc...)
	}

	strength := 0
	for i := 20; i <= len(results); i += 40 {
		val := results[i-2]
		fmt.Println(fmt.Sprintf("Result %d register is %d", i, val))
		strength += i * val
	}

	fmt.Println(fmt.Sprintf("Total signal strength: %d", strength))

	// second part
	sprite := newSprite()
	fmt.Println(len(commands))
	reg = results[0]
	cycle := 0
	for row := 0; row <= 5; row++ {
		for col := 0; col < 40; col++ {
			if sprite.covers(col) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
			reg = results[cycle]
			cycle++
			sprite.move(reg)
		}
		fmt.Println()
	}
}
