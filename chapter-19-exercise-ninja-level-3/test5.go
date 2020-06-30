package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("Saat angka %v di bagi 4 akan mengingatkan angka %v\n", i, i%4)
	}
}
