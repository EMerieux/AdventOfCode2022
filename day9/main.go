package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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
		element.PosY = leader.PosY
	} else if leader.PosX-element.PosX > 1 {
		element.PosX += 1
		element.PosY = leader.PosY
	}

	if leader.PosY-element.PosY < -1 {
		element.PosY -= 1
		element.PosX = leader.PosX
	} else if leader.PosY-element.PosY > 1 {
		element.PosY += 1
		element.PosX = leader.PosX
	}
}

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

	head := element{0, 0}
	tail := element{0, 0}
	visited := make(map[int][]int)
	visited[0] = append(visited[0], 0)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		values := strings.Split(line, " ")
		distance, err := strconv.Atoi(values[1])
		check(err)
		direction := values[0]

		for distance > 0 {
			head.move(direction)
			distance -= 1
			tail.follow(head)
			visited[tail.PosX] = append(visited[tail.PosX], tail.PosY)

			fmt.Println("Movement: " + line)
			fmt.Println("Head [" + strconv.Itoa(head.PosX) + "," + strconv.Itoa(head.PosY) + "]")
			fmt.Println("Tail [" + strconv.Itoa(tail.PosX) + "," + strconv.Itoa(tail.PosY) + "]")
		}
	}

	visitedCount := 0
	for _, ys := range visited {
		uys := make(map[int]struct{})
		for _, y := range ys {
			uys[y] = struct{}{}
		}
		visitedCount += len(uys)
	}
	fmt.Println("Tail visited: " + strconv.Itoa(visitedCount))
}
