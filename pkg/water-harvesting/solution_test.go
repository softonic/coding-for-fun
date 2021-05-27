package water_harvesting

import "testing"

type TestTuple struct {
	Input []int
	Expected int
}
func TestNoWater(t *testing.T) {
	input := []int{5,0}
	got := howMuchWater(input)
	expected := int(0)
	if got != expected {
		t.Errorf("Expected %d,int32 got %d", expected, got)
	}
}

func TestSimpleWater(t *testing.T) {
	buckets := []TestTuple{
		{ Input: []int{5,0,5}, Expected: 5},
		{ Input: []int{6,0,6}, Expected: 6},
		{ Input: []int{5,2,5}, Expected: 3},
		{ Input: []int{5,0,3}, Expected: 3},
		{ Input: []int{5,0,0,3}, Expected: 6},
		{ Input: []int{5,0,3,2,0}, Expected: 3},
		{ Input: []int{5,0,3,2,0,8}, Expected: 15},
		{ Input: []int{5,0,3,8,0,2}, Expected: 9},
		{ Input: []int{5,0,3,2,0,2}, Expected: 5},
		{ Input: []int{5,0,3,8,0,8,0,2}, Expected: 17},
		{ Input: []int{5,0,5,0,5}, Expected: 10},
	}
	for i, input := range buckets {
		got := howMuchWater(input.Input)
		expected := input.Expected
		if got != expected {
			t.Errorf("Case %d, expected %d, got %d", i, expected, got)
		}
	}
}