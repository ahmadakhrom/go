<<<<<<< HEAD
package main

import (
	"fmt"

	"github.com/ahmadakhrom/go/chapter-71-exercise-ninja-level-13/test2/quote"

	"github.com/ahmadakhrom/go/chapter-71-exercise-ninja-level-13/test2/word"
)

func main() {
	fmt.Println(word.Count(quote.AlberEinsteinSaid))           //30            //we can see how much words on quotes of Albert einstein
	for i, v := range word.UseCount(quote.AlberEinsteinSaid) { //printed array //we can print every word that in the quotes
		fmt.Println("the same word in quotes is :", v, "==>", i)
	}

	s := "andi subagja dirjo mangunharjo sujatmiko lamunglaksono"
	d := word.Count(s)
	fmt.Println(d, "---> word(s)")
	fmt.Println(len(s))
}
=======
package main

import (
	"chapter-71-exercise-ninja-level-13/test2/quote"
	"chapter-71-exercise-ninja-level-13/test2/word"
	"fmt"
)

func main() {
	fmt.Println(word.Count(quote.AlberEinsteinSaid))           //30            //we can see how much words on quotes of Albert einstein
	for i, v := range word.UseCount(quote.AlberEinsteinSaid) { //printed array //we can print every word that in the quotes
		fmt.Println("the same word in quotes is :", v, "==>", i)
	}

	s := "andi subagja dirjo mangunharjo sujatmiko lamunglaksono"
	d := word.Count(s)
	fmt.Println(d, "---> word(s)")
	fmt.Println(len(s))
}
>>>>>>> ad1ed99428e7b155e6025ad029c34a8b112d5685
