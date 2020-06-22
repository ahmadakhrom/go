package main

import "fmt"

func main() {
	var any_number1 int
	var any_number2 int
	fmt.Scanln(&any_number1)
	fmt.Println("---------")
	fmt.Scanln(&any_number2)
	some_numbs := hitung(any_number1, any_number2)
	fmt.Println(fmt.Sprintf("Average your number is : %.2f", some_numbs))
}

func hitung(number ...int) float64 {

	var total int = 0
	for _, numbers := range number {
		total += numbers //dibaca total = total + number
	}

	var avg = float64(total) / float64(len(number))
	return avg
}
