package main

import (
	"testing"
)

//get average from numbers filled
func TestGetAverage(t *testing.T) {
	val := GetAverage(41, 20)

	if val != 30.5 {
		t.Errorf("unexpected 30.5 but result was: %v", val)
	}
}
