package myfunction

import (
	"strings"
)

func JointextManualWay(text []string) string {
	s := text[0]
	for _, v := range text[1:] {
		s += " "
		s += v
	}
	return s
}

func JoinTextFastWay(text []string) string {
	return strings.Join(text, " ")
}
