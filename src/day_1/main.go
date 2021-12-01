package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readSourceFile() []string {
	var text []string

	file, err := os.Open("./src/day_1/one.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()
	return text
}

func taskOne() int {
	result := 0
	text := readSourceFile()

	for i := 1; i < len(text); i++ {
		prev, _ := strconv.Atoi(text[i-1])
		curr, _ := strconv.Atoi(text[i])
		if curr > prev {
			result += 1
		}
	}

	return result
}

func taskTwo() int {
	result := 0
	text := readSourceFile()
	preparedSlice := make([]int, len(text), len(text))

	for i, v := range text {
		preparedSlice[i], _ = strconv.Atoi(v)
	}

	previous := preparedSlice[0] + preparedSlice[1] + preparedSlice[2]
	for i := 1; i < len(preparedSlice)-2; i++ {
		current := preparedSlice[i] + preparedSlice[i+1] + preparedSlice[i+2]
		if current > previous {
			result += 1
		}
		previous = current
	}

	return result
}

func main() {
	var taskNumber int
	fmt.Scan(&taskNumber)

	if taskNumber == 1 {
		fmt.Printf("Result of task 1 of day 1 is %v", taskOne())
	}

	if taskNumber == 2 {
		fmt.Printf("Result of task 2 of day 1 is %v", taskTwo())
	}
}
