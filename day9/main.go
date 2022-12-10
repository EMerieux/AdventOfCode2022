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
	visitedStep2 := make(map[int][]int)
	visitedStep2[0] = append(visitedStep2[0], 0)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		values := strings.Split(line, " ")
		distance, err := strconv.Atoi(values[1])
		check(err)
		direction := values[0]

		for distance > 0 {
			distance -= 1
			head.move(direction)
			tail.follow(head)
			visited[tail.PosX] = append(visited[tail.PosX], tail.PosY)

			n0.move(direction)
			n1.follow(n0)
			n2.follow(n1)
			n3.follow(n2)
			n4.follow(n3)
			n5.follow(n4)
			n6.follow(n5)
			n7.follow(n6)
			n8.follow(n7)
			n9.follow(n8)
			visitedStep2[n9.PosX] = append(visitedStep2[n9.PosX], n9.PosY)

			fmt.Println("Movement: " + line)
			fmt.Println("N0 [" + strconv.Itoa(n0.PosX) + "," + strconv.Itoa(n0.PosY) + "]")
			fmt.Println("N1 [" + strconv.Itoa(n1.PosX) + "," + strconv.Itoa(n1.PosY) + "]")
			fmt.Println("N2 [" + strconv.Itoa(n2.PosX) + "," + strconv.Itoa(n2.PosY) + "]")
			fmt.Println("N3 [" + strconv.Itoa(n3.PosX) + "," + strconv.Itoa(n3.PosY) + "]")
			fmt.Println("N4 [" + strconv.Itoa(n4.PosX) + "," + strconv.Itoa(n4.PosY) + "]")
			fmt.Println("N5 [" + strconv.Itoa(n5.PosX) + "," + strconv.Itoa(n5.PosY) + "]")
			fmt.Println("N6 [" + strconv.Itoa(n6.PosX) + "," + strconv.Itoa(n6.PosY) + "]")
			fmt.Println("N7 [" + strconv.Itoa(n7.PosX) + "," + strconv.Itoa(n7.PosY) + "]")
			fmt.Println("N8 [" + strconv.Itoa(n8.PosX) + "," + strconv.Itoa(n8.PosY) + "]")
			fmt.Println("N9 [" + strconv.Itoa(n9.PosX) + "," + strconv.Itoa(n9.PosY) + "]")
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

	visitedStep2Count := 0
	for _, ys := range visitedStep2 {
		uys := make(map[int]struct{})
		for _, y := range ys {
			uys[y] = struct{}{}
		}
		visitedStep2Count += len(uys)
	}
	fmt.Println("N9 visited: " + strconv.Itoa(visitedStep2Count))
}
