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

func containsInt64(n []int64, val int64) bool {
	for _, v := range n {
		if v == val {
			return true
		}
	}

	return false
}

func containsString(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func containsLetter(s string, str string) bool {
	for _, v := range s {
		if string(v) == str {
			return true
		}
	}

	return false
}

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
