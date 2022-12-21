package seven

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/thekad/aoc/2022/pkg"
	"github.com/thekad/aoc/2022/pkg/stack"
)

const (
	dirLimit       int = 100000
	availableSpace int = 70000000
	minimumSpace   int = 30000000
)

type fsFile struct {
	size    int
	name    string
	dirpath string
}

func newFSFile(dirname string, line string) fsFile {
	s := strings.Split(line, " ")[0]
	name := strings.Join(strings.Split(line, " ")[1:], " ")

	size, _ := strconv.Atoi(s)

	return fsFile{size: size, name: name, dirpath: dirname}
}

func getContents(currdir string, output []string) ([]fsFile, []string) {
	files := []fsFile{}
	dirs := []string{}

	for _, line := range output {
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "dir ") {
			dirs = append(dirs, strings.Replace(line, "dir ", "", -1))
		} else {
			files = append(files, newFSFile(currdir, line))
		}
	}

	return files, dirs
}

// Main function of cmd "seven"
func Main(commands [][]string) {
	// first part
	fsMap := map[string]int{}
	dirStats := map[string]int{}
	dirStack := stack.NewStringStack()
	// root dir is always there
	dirStack.Push("")

	for _, command := range commands {
		cmd := command[0]
		output := command[1:]
		switch {
		case strings.HasPrefix(cmd, "ls"):
			// read the contents of the current dir
			files, _ := getContents(strings.Join(dirStack.All(), "/"), output)
			for _, file := range files {
				key := fmt.Sprintf("%s/%s", strings.Join(dirStack.All(), "/"), file.name)
				// store the file data
				fsMap[key] = file.size
				// add the size to current and parents dirs
				for i := len(dirStack.All()); i > 0; i-- {
					dirName := strings.Join(dirStack.All()[0:i], "/")
					dirStats[dirName] += file.size
				}
			}
		case cmd == "cd ..":
			dirStack.Pop()
			continue
		case strings.HasPrefix(cmd, "cd "):
			dirStack.Push(strings.Join(strings.Split(cmd, " ")[1:], " "))
			continue
		}
	}
	total := 0
	for dir, size := range dirStats {
		if size > dirLimit {
			continue
		}
		fmt.Println(fmt.Sprintf("Dir %s has size %d and can be deleted", dir, size))
		total += size
	}

	fmt.Println(fmt.Sprintf("Total size is: %d", total))

	// second part
	dirs := make(pkg.IntPairList, len(dirStats))
	i := 0
	for k, v := range dirStats {
		dirs[i] = pkg.IntPair{Key: k, Value: v}
		i++
	}

	totalUsed := dirStats[""]
	freeSpace := availableSpace - totalUsed
	target := minimumSpace - freeSpace
	sort.Sort(dirs)
	for _, v := range dirs {
		fmt.Print(fmt.Sprintf("Is %v > than %d? ", v, target))
		if v.Value > target {
			fmt.Println("Yes.")
			break
		}
		fmt.Println("Not yet.")
	}

}
