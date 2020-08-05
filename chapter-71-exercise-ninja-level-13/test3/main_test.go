package main

import (
	"chapter-71-exercise-ninja-level-13/test3/formula"
	"testing"
)

func TestCenterAverage(t *testing.T) {
	f := formula.GetCenterAverage([]int{5, 5, 5, 5}) //15/3 = 5
	if f != 5 {
		t.Errorf("unexpected %f", f)
	}
}

func BenchmarkCenterAverage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		formula.GetCenterAverage([]int{5, 5, 5, 5}) //15/3 = 5
	}
}
