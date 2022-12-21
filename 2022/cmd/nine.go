package cmd

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/thekad/aoc/2022/cmd/nine"
	"github.com/thekad/aoc/2022/pkg/file"
)

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
		nine.Main(lines)
	},
}

func init() {
	rootCmd.AddCommand(nineCmd)
}
