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

	for fileScanner.Scan() {
		line := fileScanner.Text()

		buffer := make([]byte, 0)
		for k, letter := range line {
			idx := bytes.IndexByte(buffer, byte(letter))
			if idx != -1 {
				buffer = buffer[idx+1:]
			}
			buffer = append(buffer, byte(letter))

			if len(buffer) == 14 {
				fmt.Println("Chain of X unique after character: " + strconv.Itoa(k+1))
				fmt.Print("SEQ: ")
				for _, letterInSlice := range buffer {
					fmt.Print(string(letterInSlice))
				}
				break
			}
		}
	}

	fmt.Println("")
}
