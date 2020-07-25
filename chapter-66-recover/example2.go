package main

import "fmt"

func main() {
	a()
	fmt.Println("program normally from a") //=======6

}

func a() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered in a", r) //=======5
		}
	}()

	fmt.Println("call g func..") //=======1
	g(0)
	fmt.Println("Returned normally from===== g.") // ended acces function coz in func "a" be a recover after func "g"
}

func g(x int) {
	if x > 3 {
		fmt.Println("panicking...") //========3
		panic(fmt.Sprintf("%v", x)) //stoping loop with panic
	}

	defer fmt.Println("defer in g", x) //=======4
	fmt.Println("printing in g", x)    //=======2
	g(x + 1)

}
