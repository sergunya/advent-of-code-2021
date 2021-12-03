package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	firstPart()
	secondPart()
}

func firstPart() {
	source := readSourceFile()
	length := utf8.RuneCountInString(source[0])

	gamma := make([]int, length, length)
	epsilon := make([]int, length, length)

	for i := 0; i < length; i++ {
		one, zero := 0, 0

		for j := 0; j < len(source); j++ {
			curr, _ := strconv.Atoi(strings.Split(source[j], "")[i])
			if curr == 1 {
				one++
			} else if curr == 0 {
				zero++
			}
		}
		if one > zero {
			gamma[i] = 1
			epsilon[i] = 0
		} else {
			gamma[i] = 0
			epsilon[i] = 1
		}
	}

	fmt.Printf("First part of day 3 is %v \n", getDecimal(gamma)*getDecimal(epsilon))
}

func checkPopular(s []string, row int, eq int) (int, bool) {
	one, zero := 0, 0

	for i := 0; i < len(s); i++ {
		line := strings.Split(s[i], "")
		curr, _ := strconv.Atoi(line[row])
		if curr == 1 {
			one++
		} else {
			zero++
		}
	}

	if one > zero {
		return 1, false
	} else if zero > one {
		return 0, false
	} else {
		return 1, true
	}
}

func removeLines(src *[]string, row, condition int) []string {
	res := make([]string, 0, len(*src))
	for i := 0; i < len(*src); i++ {
		line := strings.Split((*src)[i], "")
		curr, _ := strconv.Atoi(line[row])
		if curr != condition {
			res = append(res, (*src)[i])
		}
	}

	return res
}

func secondPart() {
	source := readSourceFile()
	length := utf8.RuneCountInString(source[0])

	oxygen := source
	co2 := source

	for i := 0; i < length; i++ {
		if len(oxygen) > 1 {
			popular, equal := checkPopular(oxygen, i, 1)
			if !equal {
				if popular == 1 {
					oxygen = removeLines(&oxygen, i, 0)
				} else {
					oxygen = removeLines(&oxygen, i, 1)
				}
			} else {
				oxygen = removeLines(&oxygen, i, 0)
			}

		} else {
			break
		}
	}

	for i := 0; i < length; i++ {
		if len(co2) > 1 {
			popular, equal := checkPopular(co2, i, 0)
			if !equal {
				co2 = removeLines(&co2, i, popular)
			} else {
				co2 = removeLines(&co2, i, 1)
			}
		} else {
			break
		}
	}

	ox, _ := strconv.ParseInt(oxygen[0], 2, 64)
	co, _ := strconv.ParseInt(co2[0], 2, 64)

	fmt.Printf("Second part of day 3 is %v \n", ox*co)
}

func getDecimal(sl []int) int64 {
	strRepresentation := strings.Trim(strings.Replace(fmt.Sprint(sl), " ", "", -1), "[]")
	res, _ := strconv.ParseInt(strRepresentation, 2, 64)

	return res
}

func readSourceFile() []string {
	var text []string

	file, err := os.Open("./src/day_3/source.txt")

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
