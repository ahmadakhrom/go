package main

import "fmt"

func main() {
	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	total := sum(number...)
	fmt.Println("jumlah angka dari var number :", total) //45

	even := even(sum, number...)                                    //<<== permainan congklak, sum : papan congklak, ...number : biji congklak
	fmt.Println("jumlah angka genap dalam var number 1-45 :", even) //20

	odd := odd(sum, number...)                                     //<<== permainan congklak, sum : papan congklak, ...number : biji congklak
	fmt.Println("jumlahangka ganjil dalam var number 1-45 :", odd) //25

}

func even(GetEvenNumb func(evenNumbStored ...int) int, rangeNumber ...int) int {
	//variabel "evenNumberStored" adalah tempat untuk total hasil variabel d & v yang di append
	//fmt.Printf("%T \n", GetEvenNumb)

	var d []int
	for _, v := range rangeNumber {
		if v%2 == 0 { //--------------------------maksudnya adalah menjumlahkan total angka yang kelipatan 2
			d = append(d, v)
		}
	}
	return GetEvenNumb(d...)
}

func odd(GetOddNumb func(evenNumbStored ...int) int, rangeNumber ...int) int {
	//variabel "evenNumberStored" adalah tempat untuk total hasil variabel d & v yang di append

	var d []int
	for _, v := range rangeNumber {
		if v%2 != 0 { //--------------------------maksudnya adalah menjumlahkan total angka yang buka = kelipatan 2
			d = append(d, v)
		}

	}
	return GetOddNumb(d...)
}

func sum(numb ...int) int {
	total := 0
	for _, v := range numb {
		total += v
	}
	return total
}
