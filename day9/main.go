package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type movement struct {
	Distance  int
	Direction string
}

type element struct {
	PosX int
	PosY int
}

func (element *element) move(direction string) {
	switch direction {
	case "U":
		element.PosY += 1
	case "D":
		element.PosY -= 1
	case "L":
		element.PosX -= 1
	case "R":
		element.PosX += 1
	default:
		log.Fatalf("Direction NOK")
	}
}

func (element *element) follow(leader element) {
	if leader.PosX-element.PosX < -1 {
		element.PosX -= 1
		if leader.PosY > element.PosY {
			element.PosY += 1
		} else if leader.PosY < element.PosY {
			element.PosY -= 1
		}
	} else if leader.PosX-element.PosX > 1 {
		element.PosX += 1
		if leader.PosY > element.PosY {
			element.PosY += 1
		} else if leader.PosY < element.PosY {
			element.PosY -= 1
		}
	}

	if leader.PosY-element.PosY < -1 {
		element.PosY -= 1
		if leader.PosX > element.PosX {
			element.PosX += 1
		} else if leader.PosX < element.PosX {
			element.PosX -= 1
		}
	} else if leader.PosY-element.PosY > 1 {
		element.PosY += 1
		if leader.PosX > element.PosX {
			element.PosX += 1
		} else if leader.PosX < element.PosX {
			element.PosX -= 1
		}
	}
}

func check(e error) {
	if e != nil {
		log.Fatalf("Gros naze: %v", e)
	}
}

func countVisited(visited map[int][]int) int {
	visitedCount := 0
	for _, ys := range visited {
		uys := make(map[int]struct{})
		for _, y := range ys {
			uys[y] = struct{}{}
		}
		visitedCount += len(uys)
	}
	return visitedCount
}

func part1(movements []movement) {
	head := element{0, 0}
	tail := element{0, 0}
	visited := make(map[int][]int)
	visited[0] = append(visited[0], 0)

	for _, movement := range movements {
		for movement.Distance > 0 {
			movement.Distance -= 1
			head.move(movement.Direction)
			tail.follow(head)
			visited[tail.PosX] = append(visited[tail.PosX], tail.PosY)
		}
	}

	fmt.Println("Part 1 - Tail visited: " + strconv.Itoa(countVisited(visited)))
}

func part2(movements []movement) {
	elements := make([]*element, 0)
	for i := 0; i < 10; i++ {
		elements = append(elements, &element{0, 0})
	}
	visited := make(map[int][]int)
	visited[0] = append(visited[0], 0)

	for _, movement := range movements {
		distance := movement.Distance
		for distance > 0 {
			distance -= 1
			nbElements := len(elements)
			for k, element := range elements {
				if k == 0 {
					element.move(movement.Direction)
				} else {
					element.follow(*elements[k-1])
				}
				if k == nbElements-1 {
					visited[element.PosX] = append(visited[element.PosX], element.PosY)
				}
			}
		}
	}

	fmt.Println("Part 2 - N9 visited: " + strconv.Itoa(countVisited(visited)))
}

func main() {
	f, err := os.Open("input")
	check(err)
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	movements := make([]movement, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		values := strings.Split(line, " ")
		distance, err := strconv.Atoi(values[1])
		check(err)
		movements = append(movements, movement{
			Distance:  distance,
			Direction: values[0],
		})
	}

	part1(movements)
	part2(movements)
}
