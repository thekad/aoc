package cmd

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/thekad/aoc/2022/cmd/seven"
	"github.com/thekad/aoc/2022/pkg/file"
)

var sevenCmd = &cobra.Command{
	Use:   "seven",
	Short: "Run the seventh day's exercises",
	Run: func(cmd *cobra.Command, args []string) {
		var filePath string
		var err error

		if len(args) == 0 {
			filePath, err = file.DataFilePath("day-7.txt")
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

		commands, err := file.ReadChunks(filePath, "\n$ ")
		if err != nil {
			log.Fatal(err)
		}
		seven.Main(commands)
	},
}

func init() {
	rootCmd.AddCommand(sevenCmd)
}
