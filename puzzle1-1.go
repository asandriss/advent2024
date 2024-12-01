package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
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

	var first []int
	var second []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strings := strings.Fields(line)

		if len(strings) == 0 {
			break
		}

		no1, _ := strconv.Atoi(strings[0])
		no2, _ := strconv.Atoi(strings[1])

		first = append(first, no1)
		second = append(second, no2)
	}

	slices.Sort(first)
	slices.Sort(second)

	runningTotal := 0
	for i := 0; i < len(first); i++ {
		runningTotal += int(math.Abs(float64(first[i] - second[i])))
	}

	fmt.Println("Running total", runningTotal)
}
