package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/HSLU-Student/AoC-2024/day01"
	"github.com/HSLU-Student/AoC-2024/day02"
	"github.com/HSLU-Student/AoC-2024/util"
)

const AOC = `
    ___         _________   ____ ___  __ __
   /   | ____  / ____/__ \ / __ \__ \/ // /
  / /| |/ __ \/ /    __/ // / / /_/ / // /_
 / ___ / /_/ / /___ / __// /_/ / __/__  __/
/_/  |_\____/\____//____/\____/____/ /_/   
===========================================                                           

`

func main() {
	fmt.Print(AOC)

	if len(os.Args[1:]) != 1 {
		fmt.Println("❌  Invalid number of arguments passed. Expected: 1 got:", len(os.Args[1:]))
		fmt.Println("Type `go run . <Day>` to execute the puzzle")
		return
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("❌  Invalid day argument passed.")
		return
	}

	//the puzzle registry with all puzzles
	puzzles := map[int]util.Day{
		1: day01.Day01{},
		2: day02.Day02{},
	}

	selectedPuzzle, exist := puzzles[day]

	if !exist {
		fmt.Println("⚠️  Solution not implemented.")
		return
	}

	runDay(day, selectedPuzzle)
}

func runDay(dayNo int, day util.Day) {
	input := fmt.Sprintf("day%02d", dayNo)
	solutions := []util.Solution{day.Part1(util.GetContent(input)), day.Part2(util.GetContent(input))}

	for _, s := range solutions {
		fmt.Print(s.ToString())
	}
}
