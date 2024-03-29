package cmd

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/thekad/aoc/2022/cmd/five"
	"github.com/thekad/aoc/2022/pkg/file"
)

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

		five.Main(chunks)
	},
}

func init() {
	rootCmd.AddCommand(fiveCmd)
}
