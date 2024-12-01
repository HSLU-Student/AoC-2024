package day01

import (
	"sort"
	"time"

	"github.com/HSLU-Student/AoC-2024/util"
)

type Day01 struct{}

func (d Day01) Part1(input string) util.Solution {
	starttime := time.Now()

	//create slices
	lhs := []int{}
	rhs := []int{}

	for _, line := range util.SplitContentLine(input) {
		nums := util.ParseNumbers(line)
		lhs = append(lhs, nums[0])
		rhs = append(rhs, nums[1])
	}

	//sort slices
	sort.Ints(lhs)
	sort.Ints(rhs)

	//calc diff
	res := 0
	for id, num := range lhs {
		diff := num - rhs[id]
		if diff < 0 {
			diff = diff * -1
		}
		res += diff
	}

	return util.NewSolution(res, 1, time.Since(starttime))
}

func (d Day01) Part2(input string) util.Solution {
	starttime := time.Now()

	//create map & slice
	lhs := []int{}
	rhs := make(map[int]int)

	for _, line := range util.SplitContentLine(input) {
		nums := util.ParseNumbers(line)
		lhs = append(lhs, nums[0])
		rhs[nums[1]] += 1
	}

	res := 0
	for _, num := range lhs {
		res += num * rhs[num]
	}

	return util.NewSolution(res, 2, time.Since(starttime))
}
