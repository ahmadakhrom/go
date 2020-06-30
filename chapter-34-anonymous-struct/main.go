package main

import "fmt"

func main() {
	f := struct {
		name   string
		status bool
	}{
		name:   "andy nugraha",
		status: true,
	}

	fmt.Println(f)

}
