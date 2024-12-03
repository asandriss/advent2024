package day02

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const FileName = "day02.txt"

func main() {
	file, err := os.Open(FileName)

	if err != nil {
		fmt.Printf("error opening file %s", FileName)
		return
	}

	defer file.Close()

	data, _ := parseFile(file)

	if err != nil {
		fmt.Println("parsing the file failed, exiting")
		return
	}

	// Puzzle 1
	numSafeLines := getSafeLineCount(data, false)

	fmt.Printf("total number of safe lines is [ %d ]", numSafeLines)

	// puzzle 2
	numSafeLines = getSafeLineCount(data, true)
	fmt.Printf("total number of safe lines after damppening is [%d]", numSafeLines)
}

func getSafeLineCount(data [][]int, dampen bool) int {
	result := 0

	for _, line := range data {
		isSafe := dampenSafeCheck(line, dampen)
		if isSafe {
			result++
		}
	}

	return result
}

func dampenSafeCheck(line []int, dampen bool) bool {
	if !dampen {
		return isSafeLine(line)
	}

	isSafe := isSafeLine(line)
	if isSafe {
		return true
	}

	fmt.Println("Line is not safe, try damppening it", line)

	// line is not safe, try removing with removing items
	for i := 0; i < len(line); i++ {
		subarray := make([]int, 0, len(line)-1)
		subarray = append(subarray, line[:i]...)
		subarray = append(subarray, line[i+1:]...)

		fmt.Printf("Trying Subarray (skipping index %d): %v\n", i, subarray)
		if isSafeLine(subarray) {
			fmt.Printf("Subarray was SAFE: %v\n", subarray)
			return true
		}
	}

	fmt.Println("NO SAFE Subarray FOUND")

	return false
}

func isSafeLine(line []int) bool {
	// Line is "safe" if it's all lines are either in ascending or descending order
	//   and any neibouring values differ by at least one and at most 3
	if len(line) < 2 {
		return true
	}

	first, second := line[0], line[1]

	// check first two elements to decide if the order is ascending or not
	if first == second || math.Abs(float64(first-second)) > 3.0 {
		return false
	}

	ascending := first < second

	// loop starts from the second element
	for i := 1; i < len(line); i++ {
		diff := line[i] - line[i-1]

		// check diff first
		if diff == 0 || math.Abs(float64(diff)) > 3.0 {
			return false
		}

		if (ascending && diff < 0) || (!ascending && diff > 0) {
			return false
		}
	}

	return true
}

func parseFile(file *os.File) ([][]int, error) {
	var result [][]int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		horizontalStrArr := strings.Fields(line)
		horizontalIntArr, err := parseArray(horizontalStrArr)

		if err != nil {
			return nil, fmt.Errorf("Line [%s] contains invalid characters")
		}

		result = append(result, horizontalIntArr)
	}

	return result, nil
}

func parseArray(fields []string) ([]int, error) {
	var result []int

	for _, field := range fields {
		num, err := strconv.Atoi(field)

		if err != nil {
			return nil, fmt.Errorf("%s is not a number", field)
		}

		result = append(result, num)
	}

	return result, nil
}

func getPuzzleInput(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, fmt.Errorf("could not open file %s", fileName)
	}

	defer file.Close()

	return file, nil
}
