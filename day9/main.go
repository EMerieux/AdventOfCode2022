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
	n0 := element{0, 0}
	n1 := element{0, 0}
	n2 := element{0, 0}
	n3 := element{0, 0}
	n4 := element{0, 0}
	n5 := element{0, 0}
	n6 := element{0, 0}
	n7 := element{0, 0}
	n8 := element{0, 0}
	n9 := element{0, 0}
	visited := make(map[int][]int)
	visited[0] = append(visited[0], 0)

	for _, movement := range movements {
		distance := movement.Distance
		for distance > 0 {
			distance -= 1

			n0.move(movement.Direction)
			n1.follow(n0)
			n2.follow(n1)
			n3.follow(n2)
			n4.follow(n3)
			n5.follow(n4)
			n6.follow(n5)
			n7.follow(n6)
			n8.follow(n7)
			n9.follow(n8)
			visited[n9.PosX] = append(visited[n9.PosX], n9.PosY)
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
