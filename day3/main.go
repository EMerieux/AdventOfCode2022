package main

import (
	"bufio"
	"bytes"
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

func main() {
	f, err := os.Open("input")
	check(err)
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	lowercase := []byte("abcdefghijklmnopqrstuvwxyz")
	uppercase := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	priority := 0
	groupPriority := 0
	row := 0
	currentLine := ""
	groupLines := make([]string, 0)

	for fileScanner.Scan() {
		row += 1
		line := fileScanner.Text()
		bag1 := line[:len(line)/2]
		bag2 := line[len(line)/2:]

		var commonLetter byte
		for _, letter1 := range bag1 {
			for _, letter2 := range bag2 {
				if letter1 == letter2 {
					commonLetter = byte(letter1)
				}
			}
		}

		groupLines = append(groupLines, line)

		idxLC := bytes.IndexByte(lowercase, commonLetter)
		if idxLC != -1 {
			priority += idxLC + 1
		}
		idxUC := bytes.IndexByte(uppercase, commonLetter)
		if idxUC != -1 {
			priority += idxUC + 1 + 26
		}

		if row == 1 {
			currentLine = line
		} else {
			commonLetters := ""
			for _, letter1 := range line {
				for _, letter2 := range currentLine {
					if letter1 == letter2 {
						commonLetters += string(letter1)
					}
				}
			}
			currentLine = commonLetters
		}

		if row%3 == 0 {
			idxGLC := bytes.IndexByte(lowercase, currentLine[0])
			if idxGLC != -1 {
				groupPriority += idxGLC + 1
			}
			idxGUC := bytes.IndexByte(uppercase, currentLine[0])
			if idxGUC != -1 {
				groupPriority += idxGUC + 1 + 26
			}

			row = 0
			currentLine = ""
		}
	}
	fmt.Println("Priority " + strconv.Itoa(priority))
	fmt.Println("Group Priority " + strconv.Itoa(groupPriority))
}
