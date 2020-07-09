package main

import (
	"fmt"	
	"sort"
)


type Student struct {
	Name string
	Grade int
}

type ByName []Student

//these method got from pkg.go.dev
func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Grade < a[j].Grade }

func main() {

	s1 := []Student{
		{"Andi",9},
		{"Angga",7},
		{"Budi",8},
		{"Lala",6},
	}

	
	fmt.Println("unsorted\t\t: ",s1)
	sort.Sort(ByName(s1))
	fmt.Println("sorted by increment\t: ",s1)
}