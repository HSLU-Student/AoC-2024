package day03

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/HSLU-Student/AoC-2024/util"
)

type Day03 struct{}

type Instruction struct {
	startPos int
	enable   bool
}

func (d Day03) Part1(input string) util.Solution {
	starttime := time.Now()
	res := 0
	instructionRegex := regexp.MustCompile("mul\\(\\d{1,3},\\d{1,3}\\)")
	instructions := instructionRegex.FindAllString(input, -1)
	for _, instruction := range instructions {
		numbers := getInstructionNumber(instruction)
		res += numbers[0] * numbers[1]
	}

	return util.NewSolution(res, 1, time.Since(starttime))
}

func (d Day03) Part2(input string) util.Solution {
	starttime := time.Now()
	instructionRegex := regexp.MustCompile("do(n't)?\\(\\)")
	multiplicationRegex := regexp.MustCompile("mul\\(\\d{1,3},\\d{1,3}\\)")

	instructions := instructionRegex.FindAllStringIndex(input, -1)
	multiplications := multiplicationRegex.FindAllStringIndex(input, -1)

	//create instruction map
	//bool false = disabled, bool true = enabled
	instructionMap := make(map[int]Instruction)
	for i, instruction := range instructions {
		instructionMap[i] = Instruction{instruction[1], (instruction[1] - instruction[0]) == 4}
	}

	isEnabled := true
	instructionIndex := 0
	res := 0
	i := 0

	for i < len(multiplications) {
		if multiplications[i][1] < instructionMap[instructionIndex].startPos {
			if isEnabled {
				numbers := getInstructionNumber(input[multiplications[i][0]:multiplications[i][1]])
				res += numbers[0] * numbers[1]
			}
			//multiplication processed
			i++
			continue
		}
		//next instructions enabled?
		isEnabled = instructionMap[instructionIndex].enable

		//check if has next instruction
		_, hasInstruction := instructionMap[instructionIndex+1]
		if !hasInstruction {
			break
		}

		//next instruction
		instructionIndex++

	}

	//fix up last multiplications
	if instructionMap[instructionIndex].enable {
		for j := i; j < len(multiplications); j++ {
			numbers := getInstructionNumber(input[multiplications[i][0]:multiplications[i][1]])
			res += numbers[0] * numbers[1]
		}
	}

	return util.NewSolution(res, 2, time.Since(starttime))
}

func getInstructionNumber(instruction string) []int {
	numberRegex := regexp.MustCompile("\\d+")
	numbers := numberRegex.FindAllString(instruction, -1)
	instructionNumbers := []int{}
	for _, numberStr := range numbers {
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			fmt.Println("could not convert to number")
		}
		instructionNumbers = append(instructionNumbers, number)
	}
	return instructionNumbers
}
