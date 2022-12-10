package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type State struct {
	NbCycle       int
	RegisterValue int
}

type Output struct {
	SumSignalStrength int
	CurrentColNb      int
	CurrentCrtLine    string
}

func check(e error) {
	if e != nil {
		log.Fatalf("Gros naze: %v", e)
	}
}

func executeCycle(state *State, output *Output, registerIncrement int) {
	state.NbCycle += 1

	if state.RegisterValue-1 == output.CurrentColNb || state.RegisterValue == output.CurrentColNb || state.RegisterValue+1 == output.CurrentColNb {
		output.CurrentCrtLine += "#"
	} else {
		output.CurrentCrtLine += "."
	}

	output.CurrentColNb += 1
	if output.CurrentColNb > 39 {
		output.CurrentColNb = 0
		fmt.Println(output.CurrentCrtLine)
		output.CurrentCrtLine = ""
	}

	if state.NbCycle == 20 || state.NbCycle == 60 || state.NbCycle == 100 || state.NbCycle == 140 || state.NbCycle == 180 || state.NbCycle == 220 {
		output.SumSignalStrength += state.RegisterValue * state.NbCycle
	}

	state.RegisterValue += registerIncrement
}

func main() {
	f, err := os.Open("input")
	check(err)
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	state := State{
		NbCycle:       0,
		RegisterValue: 1,
	}
	output := Output{
		SumSignalStrength: 0,
		CurrentCrtLine:    "",
	}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		values := strings.Split(line, " ")
		command := values[0]

		executeCycle(&state, &output, 0)

		if command == "addx" {
			val, err := strconv.Atoi(values[1])
			check(err)
			executeCycle(&state, &output, val)
		}
	}
	fmt.Println("Sum " + strconv.Itoa(output.SumSignalStrength))
}
