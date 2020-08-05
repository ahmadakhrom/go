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
	a := []int{2, 3, 6, 5, 6, 4, 5} // 23/5  sorting => 2345566 => 34556 => equal : 23 and len : 4
	b := []int{1, 2, 4, 6, 9}       // 12/3
	c := []int{1, 2, 3, 4, 5, 6, 7} // 20/5
	d := []int{10, 15, 10}          // 10/1
	e := [][]int{a, b, c, d}        //semua angka dijumlahkan = x, dan total digit angka - 1 = y,
	return e                        //rumus perhitungan : x/y
}
