package main

import "fmt"

func main(){

	//d := []rune("helloo") //known length of string
	//z:=0
	//s :=z <len(d)/2
	fmt.Println(ReverseString("rolod muspi merol"))
}

func ReverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
