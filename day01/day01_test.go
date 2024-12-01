package day01

import (
	"testing"

	"github.com/HSLU-Student/AoC-2024/util"
)

var INPUT = util.GetContentRoot("day01")

func TestPart1(t *testing.T) {
	expect := 2264607
	got := Day01{}.Part1(INPUT)

	if got.Result != expect {
		t.Errorf("Expected: %v, got: %v", expect, got.Result)
	}
}

func TestPart2(t *testing.T) {
	expect := 19457120
	got := Day01{}.Part2(INPUT)

	if got.Result != expect {
		t.Errorf("Expected: %v, got: %v", expect, got.Result)
	}
}
