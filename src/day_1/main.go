package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func taskOne() int {
	result := 0
	var text []string

	file, err := os.Open("./src/day_2/one.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	for i := 1; i < len(text); i++ {
		prev, _ := strconv.Atoi(text[i-1])
		curr, _ := strconv.Atoi(text[i])
		if curr > prev {
			result += 1
		}
	}

	return result

}

func main() {
	var taskNumber int
	taskNumber, err := fmt.Scan(&taskNumber)
	if err != nil {
		log.Fatalf("task number should be 1 or 2")
		return
	}

	if taskNumber == 1 {
		fmt.Printf("Result of task 1 of day 1 is %v", taskOne())
	}

}
