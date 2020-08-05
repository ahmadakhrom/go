package word

import (
	"strings"
)

func UseCount(s string) map[string]int { //we can print every word in the quotes
	as := strings.Fields(s)
	m := make(map[string]int)

	for _, v := range as {
		m[v]++
	}
	return m
}

func Count(s string) int { //we can see how much words on quotes of Albert einstein
	return len(strings.Fields(s))
}
