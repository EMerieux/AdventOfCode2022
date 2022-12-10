package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatalf("Gros naze: %v", e)
	}
}

func getSectionSlice(data string) []int {
	bounds := strings.Split(data, "-")
	lowerBound, err := strconv.Atoi(bounds[0])
	check(err)
	upperBound, err := strconv.Atoi(bounds[1])
	check(err)

	section := make([]int, 0)
	for i := 1; i <= upperBound; i++ {
		if lowerBound <= i && i <= upperBound {
			section = append(section, i)
		}
	}

	return section
}

func main() {
	f, err := os.Open("input")
	check(err)
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	nbPartiallyOverlapped := 0
	nbFullyOverlapped := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		pairs := strings.Split(line, ",")
		elf1Section := getSectionSlice(pairs[0])
		elf2Section := getSectionSlice(pairs[1])

		nbOverlaps := 0
		overlaps := false
		for _, v1 := range elf1Section {
			for _, v2 := range elf2Section {
				if v1 == v2 {
					nbOverlaps += 1
					overlaps = true
				}
			}
		}

		if len(elf1Section) == nbOverlaps || len(elf2Section) == nbOverlaps {
			nbFullyOverlapped += 1
		}
		if overlaps {
			nbPartiallyOverlapped += 1
		}
	}
	fmt.Println("Fully overlapped: " + strconv.Itoa(nbFullyOverlapped))
	fmt.Println("Partially overlapped: " + strconv.Itoa(nbPartiallyOverlapped))
}
