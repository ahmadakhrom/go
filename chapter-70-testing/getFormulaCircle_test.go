package main

import (
	"math"
	"testing"
)

var (
	shape                    = circle{20}
	diameterShouldBe float64 = shape.size / 2
	wideShoulBe      float64 = math.Phi * (math.Pow((shape.size / 2), 2))
)

//testing sum of diameter method from circle strcut
func TestDiameterAccuracy(t *testing.T) {
	t.Logf("---Diameter : %.2f", shape.diameter())

	if shape.diameter() != diameterShouldBe {
		t.Errorf("false! not: %.2f, but diameter should be : %.2f", shape.diameter(), diameterShouldBe)
	}
}

//testing som of wide method from circle strcut
func TestWideAccuracy(t *testing.T) {
	t.Logf("---Wide : %.2f", shape.wide())

	if shape.wide() != wideShoulBe {
		t.Errorf("false! not: %.2f, but wide should be : %.2f", shape.wide(), wideShoulBe)
	}
}

//testing get sum, from answer struct
func TestGetSum(t *testing.T) {
	person := Answer{10, []int{2, 3, 4, 1}}

	s := GetSum(person.result)
	if s != person.Guess {
		t.Error("your guess is false:", person.Guess, "should be:", s)
	}
}
