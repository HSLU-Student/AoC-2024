package dayXX

import (
	"testing"

	"github.com/HSLU-Student/AoC-2024/util"
)

var INPUT = util.GetContentRoot("dayXX")

func TestPart1(t *testing.T) {
	expect := 0
	got := DayXX{}.Part1(INPUT)

	if got.Result != expect {
		t.Errorf("Expected: %v, got: %v", expect, got.Result)
	}
}

func TestPart2(t *testing.T) {
	expect := 0
	got := DayXX{}.Part2(INPUT)

	if got.Result != expect {
		t.Errorf("Expected: %v, got: %v", expect, got.Result)
	}
}
