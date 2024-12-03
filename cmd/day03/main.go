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
	runningSum := 0

	for _, el := range fullContent {
		commands, _ := splitByCommand(el)

		validContent, _ := regexFilter(commands, `mul\(\s*\d+\s*,\s*\d+\s*\)`)
		for _, el := range validContent {
			mul, _ := getMultiplication(el)
			runningSum += mul
		}

	}
	fmt.Println("Total multiplications is ", runningSum)
}

func processCommand(cmd string) {

}

func splitByCommand(text string) ([]string, error) {
	pattern := `(do\(\)|don't\(\))`

	re, err := regexp.Compile(pattern)

	if err != nil {
		return nil, fmt.Errorf("error while building regex %w", err)
	}

	matches := re.FindAllStringIndex(text, -1)
	var parts []string

	prev := 0
	for _, match := range matches {
		if len(match) < 2 {
			continue // ignore empty
		}

		start, end := match[0], match[1]

		if start > prev && text[start:end] != "don't()" {
			parts = append(parts, text[prev:start])
		}

		// parts = append(parts, text[start:end])
		prev = end

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
