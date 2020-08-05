package main

import (
	"chapter-71-exercise-ninja-level-13/test3/formula"
	"fmt"
)

func main() {
	res := dataNumber()
	for _, v := range res {
		fmt.Println(formula.GetCenterAverage(v))
	}
}

func dataNumber() [][]int {
	a := []int{2, 3, 6, 5, 6, 4, 5} // 3+4+5+5+6+6 / 6 = 4,83
	b := []int{4, 5, 6}             // 11/2 5.5
	c := []int{1, 2, 3, 4, 5, 6, 7} // 27/6 = 4.5
	d := []int{10, 15, 10}          // 25/2 = 12.5
	e := [][]int{a, b, c, d}        //semua angka dijumlahkan = x, dan total digit angka - 1 = y,
	return e                        //rumus perhitungan : x/y
}
