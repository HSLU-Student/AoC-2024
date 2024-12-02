package day02

import (
	"slices"
	"time"

	"github.com/HSLU-Student/AoC-2024/util"
)

type Day02 struct{}

func (d Day02) Part1(input string) util.Solution {
	starttime := time.Now()
	reports := []string{}
	reports = append(reports, util.SplitContentLine(input)...)

	res := 0
	for _, report := range reports {
		numbers := util.ParseNumbers(report)
		valid := isValidReport(numbers)
		if valid {
			res += 1
		}
	}
	return util.NewSolution(res, 1, time.Since(starttime))
}

func (d Day02) Part2(input string) util.Solution {
	starttime := time.Now()
	reports := []string{}
	reports = append(reports, util.SplitContentLine(input)...)

	res := 0
	for _, report := range reports {
		numbers := util.ParseNumbers(report)
		for i := 0; i < len(numbers); i++ {
			slicedReport := slices.Concat(numbers[:i], numbers[i+1:])
			valid := isValidReport(slicedReport)
			if valid {
				res += 1
				break
			}

		}
	}

	return util.NewSolution(res, 2, time.Since(starttime))
}

func isValidReport(report []int) bool {
	valid := true
	lastdiff := 0
	for i := range len(report) - 1 {
		valid, lastdiff = isValidStep(report[i], report[i+1], lastdiff)
		if !valid {
			break
		}
	}
	return valid
}

func isValidStep(left int, right int, lastdiff int) (bool, int) {
	diff := left - right
	if diff > 3 || diff < -3 || diff == 0 {
		return false, 0
	}

	if lastdiff < 0 && diff > 0 {
		return false, 0
	}

	if lastdiff > 0 && diff < 0 {
		return false, 0
	}
	return true, diff
}
