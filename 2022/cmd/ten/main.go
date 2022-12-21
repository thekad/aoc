package ten

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// Main method for cmd "ten"
func Main(lines []string) {
	spew.Dump(lines)
	fmt.Println("done")
}
