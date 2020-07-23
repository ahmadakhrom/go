package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	namefile := "logs.txt"
	foo(namefile)
}

func foo(fileName string) {
	f, err := os.Create(fileName)
	defer f.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	r := strings.NewReader("Hello world")
	io.Copy(f, r)
}
