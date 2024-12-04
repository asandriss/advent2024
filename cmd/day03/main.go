package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const FileName = "day03.txt"

func main() {
	file, err := os.Open(FileName)
	if err != nil {
		fmt.Println("error opening file", err)
		return
	}

	defer file.Close()

	fullContent := getFileContent(file)
	runningSum := solve(fullContent)

	fmt.Println("Total multiplications is ", runningSum)
}

func solve(fullContent []string) int {

	runningSum := 0

	for _, el := range fullContent {
		doCommands, _ := filterDosAndDonts(el)
		// fmt.Println("filtered do commands ", doCommands)

		mulCommands, _ := regexFilter(doCommands, `mul\(\s*\d+\s*,\s*\d+\s*\)`)
		for _, mulCmd := range mulCommands {
			mul, _ := getMultiplication(mulCmd)
			runningSum += mul
		}

	}

	return runningSum
}

func filterDosAndDonts(text string) ([]string, error) {
	pattern := `(do\(\)|don't\(\))`

	re, err := regexp.Compile(pattern)

	if err != nil {
		return nil, fmt.Errorf("error while building regex %w", err)
	}

	matches := re.FindAllStringIndex(text, -1)

	var parts []string
	inDoBlock := true // assume the do() block for first line

	prev := 0
	for _, match := range matches {
		if len(match) < 2 {
			continue // ignore empty
		}

		start, end := match[0], match[1]

		if inDoBlock && start > prev {
			matchedText := text[prev:start]

			parts = append(parts, matchedText)
		}
		prev = end

		inDoBlock = (text[start:end] == "do()")
	}
	if inDoBlock && prev < len(text) {
		remainingText := text[prev:]
		parts = append(parts, remainingText)
	}
	return parts, nil
}

func getMultiplication(text string) (int, error) {
	re, err := regexp.Compile(`\d+`)

	if err != nil {
		return 0, fmt.Errorf("error while building regex %w", err)
	}

	matches := re.FindAllString(text, 2)
	first, _ := strconv.Atoi(matches[0])
	second, _ := strconv.Atoi(matches[1])

	return first * second, nil
}

func getFileContent(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var content []string

	for scanner.Scan() {
		line := scanner.Text()
		content = append(content, line)
	}

	return content
}

func regexFilter(content []string, pattern string) ([]string, error) {
	re, err := regexp.Compile(pattern)

	if err != nil {
		return nil, fmt.Errorf("error while building regex [%s] %w", pattern, err)
	}

	var results []string

	for _, el := range content {
		if len(el) < 1 {
			continue // ignore empty lines
		}

		// fmt.Println("Applying regex to string", el)
		matches := re.FindAllString(el, -1)
		// fmt.Println("matches ", matches)
		results = append(results, matches...)
	}

	return results, nil
}
