package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	//src
	plaintText := "iloveyoubut250294v3"

	//generate password
	passGenerated, err := bcrypt.GenerateFromPassword([]byte(plaintText), bcrypt.MinCost)
	if err != nil {
		fmt.Println("error generating")
		return //----------------------------- return aim to stop the program
	}
	fmt.Println("your password generated :", string(passGenerated))

	//verify password
	filledText := "iloveyoubut250294v3"
	err = bcrypt.CompareHashAndPassword(passGenerated, []byte(filledText))
	if err != nil {
		fmt.Println("wrong password")
		return
	}
	fmt.Println("password verified")

}
