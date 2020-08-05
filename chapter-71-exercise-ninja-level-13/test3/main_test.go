package main

import (
	"testing"

	"github.com/ahmadakhrom/go/chapter-71-exercise-ninja-level-13/test3/formula"
)

func TestCenterAverage(t *testing.T) {
	type test struct {
		data   []int
		answer float64
	}

	tests := []test{
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{2, 3, 4, 5, 6}, 4},
		{[]int{2, 5, 6, 7, 7}, 6},
	}

	for _, v := range tests {
		s := formula.GetCenterAverage(v.data)
		if s != v.answer {
			t.Errorf("error! expected %v, got : %v", s, v.answer)
		}
	}
}

func BenchmarkCenterAverage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		formula.GetCenterAverage([]int{5, 3, 4, 5, 3, 2, 1, 6, 8, 25000}) //15/3 = 5
	}
}

func ExampleCenterAverage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		formula.GetCenterAverage([]int{1, 2, 3, 4, 5})
		//Output :
		//3
	}
}
