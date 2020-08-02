package benchmark-example
//package main

import (
	"fmt"
	"strings"
)

// func main() {
// 	text := "where we go today guys"
// 	d := strings.Split(text, " ")
// 	j := JoinManual(d)
// 	f := JoinText(d)

// 	fmt.Println(d) //make strings to array
// 	fmt.Println(j) //make one string
// 	fmt.Println(f) //make one string
// }

func JoinManual(text []string) string {
	s := text[0]
	for _, v := range text[1:] {
		s += " "
		s += v
	}
	return s
}

func JoinText(text []string) string {
	return strings.Join(text, " ")
}
