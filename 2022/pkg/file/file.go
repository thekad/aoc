package file

import (
	"bufio"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// DataFilePath returns the path to the data file
func DataFilePath(f string) (string, error) {
	mydir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	fPath := path.Join(mydir, "data", f)

	return fPath, nil
}

// ReadLines reads the given file and returns an array of strings
func ReadLines(filePath string) ([]string, error) {
	var lines []string

	f, err := os.Open(filePath)
	if err != nil {
		return lines, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, err
}

// ReadChunks reads the given file and splits using a separator, and returns an array of array of strings
func ReadChunks(filePath, separator string) ([][]string, error) {
	var chunks [][]string
	var content []byte

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return chunks, err
	}

	for _, chunk := range strings.Split(string(content), separator) {
		clean := strings.Split(chunk, "\n")
		chunks = append(chunks, clean)
	}

	return chunks, nil
}
