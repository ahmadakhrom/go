package main

import "fmt"

//rumus function
//func (r receiver) identifier(argument) (return(s)) { code here... }

//just print out string
func SayWelcome() {
	fmt.Println("Your Welcome.. ")
}

//passing value
func SayHello(s string) {
	fmt.Println("Hallo, ", s)
}

//passing via return % store to variable then print
func SayKuddos(name string) string {
	return fmt.Sprint("kuddos for, ", name)
}

//func multiple return
func sayGoodbye(words1, words2 string) (string, bool) {
	words := fmt.Sprint("good bye and good luck.. ", words1, " ", words2)
	err := true
	return words, err
}

func main() {
	SayWelcome()
	SayHello("john_doe")

	x := SayKuddos("andy")
	fmt.Println(x)

	y, z := sayGoodbye("budi", "budiman")
	fmt.Println(y)
	fmt.Println(z)
}
