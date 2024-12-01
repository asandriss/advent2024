package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("puzzle1-1.input.txt")
	if err != nil {
		fmt.Println("error opening file", err)
		return
	}

	defer file.Close()

	first, second, err := parseFile(file)

	if err != nil {
		fmt.Println("Error parsing the file:", err)
		return
	}

	sort.Ints(first)
	sort.Ints(second)

	runningTotal := 0
	for i := 0; i < len(first); i++ {
		runningTotal += int(math.Abs(float64(first[i] - second[i])))
	}

	fmt.Println("Running total", runningTotal)
}

func parseFile(file *os.File) ([]int, []int, error) {
	var first, second []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) < 2 {
			continue
			// return nil, nil, fmt.Errorf("invalid line format: %s", line)
		}

		left, err1 := strconv.Atoi(fields[0])
		right, err2 := strconv.Atoi(fields[1])

		if err1 != nil || err2 != nil {
			return nil, nil, fmt.Errorf("line %s contains an invalid number", line)
		}

		first = append(first, left)
		second = append(second, right)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error while reading the file %w", err)
	}

	return first, second, nil
}
