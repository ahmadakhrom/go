package main

import "testing"
import "github.com/ahmadakhrom/go"

//testing get total, from gettotal func
func TestGetTotal(t *testing.T) {
	person := []Answer{
		{15, []int{1, 2, 3, 4, 5}},
		{15, []int{1, 2, 3, 4, 5}},
		{15, []int{1, 2, 3, 4, 5}},
		{15, []int{1, 2, 3, 4, 5}},
		{15, []int{1, 2, 3, 4, 5}},
	}
	for _, v := range person {
		s := GetTotal(v.result...)
		if s != v.Guess {
			t.Errorf("your guess is false: %v, should be: %v", v.Guess, s)
		}
	}
}
