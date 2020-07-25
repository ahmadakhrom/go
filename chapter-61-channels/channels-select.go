package main

import "fmt"

func main() {

	fmt.Println("starting..")

	en := make(chan int)
	od := make(chan int)
	qu := make(chan int)
	num := 500

	//sending channels
	go Dice(en, od, qu, num)

	//receive channels
	DiceResult(en, od, qu)

	fmt.Println("done!")

}

//sending channels
func Dice(e, o, q chan int, num int) {
	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

//receive channels
func DiceResult(e, o, q chan int) {
	for {
		select {
		case val := <-e:
			fmt.Println("from even--", val)
		case val := <-o:
			fmt.Println("from  odd--", val)
		case val := <-q:
			fmt.Println("from quit--", val)
			close(e)
			close(o)
			close(q)
			return
		}
	}
}
