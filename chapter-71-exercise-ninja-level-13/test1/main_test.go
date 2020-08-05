package main

import (
	"chapter-71-exercise-ninja-level-13/test1/chicken"
	"testing"
)

func TestChicken(t *testing.T) {
	v := chicken.Years(10)
	if v != 70 {
		t.Errorf("error! unexpected : %d", v)
	}
}

func TestChicken2(t *testing.T) {
	v := chicken.YearsTwo(20)
	if v != 140 {
		t.Errorf("error! unexpected : %d", v)
	}
}

func BenchmarkChicken(b *testing.B) {
	for i := 0; i < b.N; i++ {
		chicken.Years(10)
	}
}

func BenchmarkChicken2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		chicken.YearsTwo(10)
	}
}
