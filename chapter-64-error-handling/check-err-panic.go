package main

import "fmt"

func main() {
	var answ1, answ2, answ3 string

	fmt.Print("Name: ")
	_, err := fmt.Scanln(&answ1)
	if err != nil {
		panic(err) //equivalent with os.Exit(2)
	}

	fmt.Print("Name: ")
	_, err = fmt.Scanln(&answ2)
	if err != nil {
		panic(err) //equivalent with os.Exit(2)
	}

	fmt.Print("Name: ")
	_, err = fmt.Scanln(&answ3)
	if err != nil {
		panic(err) //equivalent with os.Exit(2)
	}

	fmt.Println("Hallo ", answ1, answ2, "and", answ3)
}
