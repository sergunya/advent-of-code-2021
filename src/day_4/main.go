package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getEmptyMap() [][]string {
	res := make([][]string, 0, 1000)

	for j := 0; j < 5; j++ {
		res = append(res, make([]string, 5, 1000))
	}

	return res
}

func readMaps() [][][]string {
	var text []string
	countLines := 0
	result := make([][][]string, 0, 1000)

	file, err := os.Open("./src/day_4/source.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	currentMap := getEmptyMap()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text = strings.Fields(scanner.Text())
		if countLines < 5 {
			currentMap[countLines] = text
			countLines++
		} else {
			result = append(result, currentMap)
			countLines = 0
			currentMap = getEmptyMap()
		}
	}

	file.Close()
	return result
}

func readNumbers() []string {
	file, err := os.Open("./src/day_4/source2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	numbers := scanner.Text()
	return strings.Split(numbers, ",")
}

func generateEmptyMaps(size int) [][][]int {
	res := make([][][]int, 0, 1000)

	for i := 0; i < size; i++ {
		tmp := make([][]int, 0, 1000)
		for j := 0; j < 5; j++ {
			tmp = append(tmp, make([]int, 5, 1000))
		}

		res = append(res, tmp)
	}

	return res
}

func play(untilLast bool) {
	maps := readMaps()
	resultMaps := generateEmptyMaps(len(maps))
	numbers := readNumbers()
	winners := make([]bool, len(maps), len(maps))

	for i, v := range numbers {
		for j, matrix := range maps {
			if i > 3 {
				win, res := markNumber(&resultMaps[j], &matrix, v, true)
				if win {
					if untilLast {
						markWinner(&winners, j)
						nonWinnersNum := getNonWinnersNumber(&winners)

						if nonWinnersNum == 0 {
							fmt.Printf("The winner of second part is %d\n", res)
							return
						}
					} else {
						fmt.Printf("The winner of first part is %d\n", res)
						return
					}

				}
			} else {
				markNumber(&resultMaps[j], &matrix, v, false)
			}
		}
	}

	fmt.Println("Couldn't determine the winner!")
}

func markWinner(winners *[]bool, bordNumber int) {
	if !(*winners)[bordNumber] {
		(*winners)[bordNumber] = true
	}
}

func getNonWinnersNumber(winners *[]bool) int {
	count := 0

	for _, v := range *winners {
		if !v {
			count += 1
		}
	}

	return count
}

func markNumber(resMatrix *[][]int, origMatrix *[][]string, v string, checkWin bool) (bool, int) {
	for i, xValue := range *origMatrix {
		for j, yValue := range xValue {
			if yValue == v {
				(*resMatrix)[i][j] = 1

				if checkWin == true {
					if checkRow(resMatrix, i) || checkCol(resMatrix, j) {
						sum := sumUnmarkedNumbers(resMatrix, origMatrix)
						val, _ := strconv.Atoi(v)

						return true, val * sum
					}
				}
			}
		}
	}

	return false, -1
}

func sumUnmarkedNumbers(resMatrix *[][]int, matrix *[][]string) int {
	var sum = 0
	for i, x := range *resMatrix {
		for j, y := range x {
			if y == 0 {
				val, _ := strconv.Atoi((*matrix)[i][j])
				sum += val
			}
		}
	}

	return sum
}

func checkCol(matrix *[][]int, col int) bool {
	for i := 0; i < len(*matrix); i++ {
		if (*matrix)[i][col] != 1 {
			return false
		}
	}

	return true
}

func checkRow(resMatrix *[][]int, row int) bool {
	for _, v := range (*resMatrix)[row] {
		if v != 1 {
			return false
		}
	}

	return true
}

func main() {
	play(false)
	play(true)
}
