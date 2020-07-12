package main

import (
	"fmt"
	"sort"
)

type Teacher struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByLast []Teacher

//these method got from pkg.go.dev
func (a ByLast) Len() int           { return len(a) }
func (a ByLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLast) Less(i, j int) bool { return a[i].Last < a[j].Last }

type ByAge []Teacher

//these method got from pkg.go.dev
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {

	dataTeacher := []Teacher{
		{
			"Miss",
			"MennyPenny",
			27,
			[]string{
				"Oh, James. You didn't.",
				"Would you like me to take care of that for you, James?",
				"I would really prefer to be a secret agent myself."},
		},
		{
			"James",
			"Bond",
			32,
			[]string{
				"James, it is soo good to see you",
				"Dear God, what has James done now?",
				"I would really prefer to be a secret agent myself."},
		},
		{
			"Andi",
			"Nugraha",
			28,
			[]string{
				"James, it is soo good to see you",
				"Would you like me to take care of that for you, James?",
				"I would really prefer to be a secret agent myself."},
		},
	}

	//sort data by sayings
	sort.Sort(ByLast(dataTeacher))
	for _, v := range dataTeacher {
		fmt.Println(v.First, v.Last, v.Age)
		sort.Strings(v.Sayings)
		for _, val := range v.Sayings {
			fmt.Println("\t", val)
		}
	}

	//sort data by age
	sort.Sort(ByAge(dataTeacher))
	for _, v := range dataTeacher {
		fmt.Println(v.First, v.Last, v.Age)
		sort.Strings(v.Sayings)
		for _, val := range v.Sayings {
			fmt.Println("\t\t", val)
		}
	}
}
