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

func main() {
	f, err := os.Open("input")
	check(err)
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	stacks := make(map[int][]string, 0)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if len(line) > 0 {
			if line[:4] == "move" {
				elements := strings.Split(line, " ")
				qtty, err := strconv.Atoi(elements[1])
				check(err)
				src, err := strconv.Atoi(elements[3])
				check(err)
				dest, err := strconv.Atoi(elements[5])
				check(err)

				elts, newSrcStack := stacks[src][0:qtty], stacks[src][qtty:]
				stacks[src] = newSrcStack
				newDestStack := make([]string, 0)
				for _, v := range elts {
					newDestStack = append(newDestStack, v)
				}
				for _, v := range stacks[dest] {
					newDestStack = append(newDestStack, v)
				}
				stacks[dest] = newDestStack

				//for i := 1; i <= qtty; i++ {
				//	elt, newStack := stacks[src][0], stacks[src][1:]
				//	stacks[src] = newStack
				//	stacks[dest] = append([]string{elt}, stacks[dest]...)
				//}
			} else if string(line[1]) != "1" {
				nbStacks := (len(line) + 1) / 4
				for i := 1; i <= nbStacks; i++ {
					letter := string(line[((i-1)*4)+1])
					if letter != " " {
						stacks[i] = append(stacks[i], letter)
					}
				}
			}
		}
	}

	for i := 1; i < len(stacks)+1; i++ {
		fmt.Print(stacks[i][0])
	}
	fmt.Println("")
}
