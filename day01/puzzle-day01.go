package day01

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const FileName = "puzzle1-1.input.txt"

func main() {
	file, err := os.Open(FileName)
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

	// solve puzzle 1
	runningTotal(first, second)

	// solve puzzle 2
	similarityScore(first, second)
}

func similarityScore(first, second []int) {
	counts := countElements(second)
	total := 0

	for _, elem := range first {
		total += elem * counts[elem]
	}

	fmt.Println("Similarity score is", total)
}

func countElements(arr []int) map[int]int {
	result := make(map[int]int)

	for _, elem := range arr {
		if _, exists := result[elem]; exists {
			result[elem]++
		} else {
			result[elem] = 1
		}
	}

	return result
}

func runningTotal(first, second []int) {
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
