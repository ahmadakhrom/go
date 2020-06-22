package main

import "fmt"

func main(){
	some_numbs := hitung(5,5,5,5,5,5)
	fmt.Println(fmt.Sprintf("Average your number is : %.2f",some_numbs))
}


func hitung(number ...int) float64{

var total int = 0
for _, numbers := range number {		
	total += numbers	 //dibaca total = total + number
	}
	
	var avg = float64(total) / float64(len(number))
	return avg
}
