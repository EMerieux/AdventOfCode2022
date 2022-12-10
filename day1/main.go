package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	currentElfCalories := 0
	caloriesPerElf := make([]int, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.Trim(line, " ") == "" {
			caloriesPerElf = append(caloriesPerElf, currentElfCalories)
			currentElfCalories = 0
		} else {
			calories, err := strconv.Atoi(line)
			check(err)
			currentElfCalories += calories
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(caloriesPerElf)))

	fmt.Println("Elf with most calories: " + strconv.Itoa(caloriesPerElf[0]))
	fmt.Println("Top 3 calories: " + strconv.Itoa(caloriesPerElf[0]+caloriesPerElf[1]+caloriesPerElf[2]))
}
