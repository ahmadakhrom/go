package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {

	//how to print with os.stdout package
	fmt.Fprintln(os.Stdout, "hello world..")
	io.WriteString(os.Stdout, "hello world.. \n")
	fmt.Printf("%q \n", `"here \ we go.."`)


}
