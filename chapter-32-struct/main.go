package main

import "fmt"

type book struct {
	page  int
	about string
}

func main() {

	learning := book{
		page:  120,
		about: "code golang",
	}

	guide := book{
		page:  100,
		about: "cook a cake",
	}

	fmt.Println(learning)
	fmt.Println(guide)

	fmt.Println(learning.about, "\t", learning.page)
	fmt.Println(guide.about, "\t", guide.page)

}
