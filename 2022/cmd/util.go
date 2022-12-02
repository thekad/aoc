package cmd

import (
	"bufio"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func dataPath(f string) (string, error) {
	mydir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	fPath := path.Join(mydir, "data", f)

	return fPath, nil
}

func readLines(filePath string) ([]string, error) {
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

func readChunks(filePath string) ([][]string, error) {
	var chunks [][]string
	var content []byte

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return chunks, err
	}

	for _, chunk := range strings.Split(string(content), "\n\n") {
		clean := strings.Split(chunk, "\n")
		chunks = append(chunks, clean)
	}

	return chunks, nil
}
