package main

import "fmt"

type Person struct {
	First string
	Last string
}

type Say struct {
	Person
	Said string
}

func (p Person) introduce()  {
	fmt.Printf("Hello i'm %s %s \n",p.First,p.Last)
}

func (ps Say) introduce()  {
	fmt.Printf("Hello i'm %s %s, '%s' \n",ps.First,ps.Last, ps.Said)
}

type human interface {
	introduce()
}

func sayWhat(h human)  {
	h.introduce()
}

func main()  {
	Angga := Person{
		"Ahmad",
		"Maulana",
	}
	fmt.Println(Angga)

	girl := Say{
		Person{
			"Baby",
			"Young",
		},
		"Glad to meet you here",
	}
	//fmt.Println(girl)
	//girl.introduce()

	sayWhat(Angga)
	sayWhat(girl)

}