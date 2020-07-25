package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	namefile := "logs.txt"
	foo(namefile)
}

func foo(fileName string) {
	file, err := os.Open(fileName)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(resp))
}
