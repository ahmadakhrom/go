package main

import (
	"testing"

	"chapter-71-exercise-ninja-level-13/test2/quote"
	"chapter-71-exercise-ninja-level-13/test2/word"
)

func TestUseCount(t *testing.T) {
	m := word.UseCount("do re re mi mi fa fa fa")
	for k, v := range m {
		switch k {
		case "do":
			if v != 1 {
				t.Errorf("error! unexpected %d want 1", v)
			}
		case "re":
			if v != 2 {
				t.Errorf("error! unexpected %d want 1", v)
			}
		case "mi":
			if v != 2 {
				t.Errorf("error! unexpected %d want 1", v)
			}
		case "fa":
			if v != 3 {
				t.Errorf("error! unexpected %d want 1", v)
			}
		}
	}
}

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
