package day02

import (
	"testing"

	"github.com/HSLU-Student/AoC-2024/util"
)

var INPUT = util.GetContentRoot("day02")

func TestPart1(t *testing.T) {
	expect := 585
	got := Day02{}.Part1(INPUT)

	if got.Result != expect {
		t.Errorf("Expected: %v, got: %v", expect, got.Result)
	}
}

func TestPart2(t *testing.T) {
	expect := 626
	got := Day02{}.Part2(INPUT)

	if got.Result != expect {
		t.Errorf("Expected: %v, got: %v", expect, got.Result)
	}
}
