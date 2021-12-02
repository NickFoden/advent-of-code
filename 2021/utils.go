package adventOfCode

import (
	"bufio"
	"os"
	"strconv"
)

func readLines(path string) ([]int64, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		lines = append(lines, val)
	}
	return lines, scanner.Err()
}

func readStringLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
