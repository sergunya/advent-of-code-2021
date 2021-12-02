package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readSourceFile() []string {
	var text []string

	file, err := os.Open("./src/day_2/source.txt")

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

func main() {
	depth, forward := 0, 0
	source := readSourceFile()

	for _, v := range source {
		command := strings.Split(v, " ")
		switch command[0] {
		case "forward":
			p := GetValue(command[1])
			forward += p
		case "down":
			p := GetValue(command[1])
			depth += p
		case "up":
			p := GetValue(command[1])
			depth -= p

		}
	}

	fmt.Printf("Result is %d", depth*forward)

}

func GetValue(c string) int {
	p, err := strconv.Atoi(c)

	if err != nil {
		log.Fatalf("cannot parse value of forward")
	}
	return p
}
