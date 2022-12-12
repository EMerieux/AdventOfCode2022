package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		log.Fatalf("Gros naze: %v", e)
	}
}

func computeVisibleTrees(visibleTrees map[string]struct{}, rows map[int][]int, isCol bool, isReverse bool) {
	for k, row := range rows {
		maxHeight := -1
		rowL := len(row)
		for j, height := range row {
			if height > maxHeight {
				maxHeight = height
				var x int
				var y int
				if !isReverse && !isCol {
					x = j + 1
					y = k
				} else if isReverse && !isCol {
					x = rowL - j
					y = k
				} else if !isReverse && isCol {
					x = k
					y = j + 1
				} else if isReverse && isCol {
					x = k
					y = rowL - j
				}
				visibleTrees[strconv.Itoa(x)+","+strconv.Itoa(y)] = struct{}{}
			}
		}
	}
}

func main() {
	f, err := os.Open("input")
	check(err)
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	visibleTrees := make(map[string]struct{})

	rows := make(map[int][]int, 0)
	reverseRows := make(map[int][]int, 0)
	columns := make(map[int][]int, 0)
	reverseColumns := make(map[int][]int, 0)

	lineNb := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()

		lineNb += 1
		for k, v := range line {
			height, err := strconv.Atoi(string(v))
			check(err)

			rows[lineNb] = append(rows[lineNb], height)
			reverseRows[lineNb] = append([]int{height}, reverseRows[lineNb]...)
			columns[k+1] = append(columns[k+1], height)
			reverseColumns[k+1] = append([]int{height}, reverseColumns[k+1]...)
		}
	}

	computeVisibleTrees(visibleTrees, rows, false, false)
	computeVisibleTrees(visibleTrees, reverseRows, false, true)
	computeVisibleTrees(visibleTrees, columns, true, false)
	computeVisibleTrees(visibleTrees, reverseColumns, true, true)

	fmt.Println("Nb visible trees: " + strconv.Itoa(len(visibleTrees)))

	maxScore := 0
	rowL := len(rows)
	colL := len(columns)
	for rowNb, heights := range rows {
		for colNb, height := range heights {
			rightCount := 0
			walk := colNb + 1
			for walk < rowL && height > rows[rowNb][walk] {
				walk += 1
				rightCount += 1
			}
			if walk < rowL {
				rightCount += 1
			}
			leftCount := 0
			walk = colNb - 1
			for walk >= 0 && height > rows[rowNb][walk] {
				walk -= 1
				leftCount += 1
			}
			if walk >= 0 {
				leftCount += 1
			}
			upCount := 0
			walk = rowNb - 1
			for walk > 0 && height > rows[walk][colNb] {
				walk -= 1
				upCount += 1
			}
			if walk > 0 {
				upCount += 1
			}
			downCount := 0
			walk = rowNb + 1
			for walk < colL+1 && height > rows[walk][colNb] {
				walk += 1
				downCount += 1
			}
			if walk < colL+1 {
				downCount += 1
			}

			score := leftCount * rightCount * upCount * downCount
			if score > maxScore {
				maxScore = score
			}

			//fmt.Print(rowNb)
			//fmt.Print("/")
			//fmt.Print(colNb)
			//fmt.Print(": ")
			//fmt.Print(leftCount)
			//fmt.Print(" - ")
			//fmt.Print(upCount)
			//fmt.Print(" - ")
			//fmt.Print(rightCount)
			//fmt.Print(" - ")
			//fmt.Print(downCount)
			//fmt.Print(": ")
			//fmt.Print(leftCount * rightCount * upCount * downCount)
			//fmt.Println("")
		}
	}
	fmt.Println("Max score: " + strconv.Itoa(maxScore))
}
