package main
import("fmt")

func main() {

	ch := make(chan int)

	ch <- 100 //-------------channel must be wrapped on a func / variable / closure
fmt.Println(<-ch)


}


