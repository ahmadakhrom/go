package main

import (
	"fmt"
	"sort"
)

func main() {
	//sort int
	s := []int{2, 4, 2, 6, 8, 9, 1, 90, 34, 54, 4}
	sort.Ints(s)
	// fmt.Println(s)

	//sort string
	g := []string{"guava", "melon", "grape", "banana", "water melon"}
	sort.Strings(g)
	// fmt.Println(g)

	//are string / int sorted ?
	fmt.Println(sort.IntsAreSorted(s))
	fmt.Println(sort.StringsAreSorted(g))

}