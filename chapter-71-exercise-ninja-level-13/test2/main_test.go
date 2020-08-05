package main

import (
	"chapter-71-exercise-ninja-level-13/test2/quote"
	"chapter-71-exercise-ninja-level-13/test2/word"
	"testing"
)

func TestCount(t *testing.T) {
	d := word.Count(quote.AlberEinsteinSaid)
	if d != 30 {
		t.Errorf("error! unexpected : %d", d)
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		word.Count(quote.AlberEinsteinSaid)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		word.UseCount("andi subagja dirjo mangunharjo sujatmiko lamunglaksono")
	}
}
