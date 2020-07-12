package main

import "fmt"

func main() {

	fmt.Println("starting..")

	en := make(chan int)
	od := make(chan int)
	qu := make(chan bool)
	num := 15

	//sending channels
	go Dice(en, od, qu, num)

	//receive channels
	DiceResult(en, od, qu)

	fmt.Println("done!")

}

//sending channels
func Dice(e, o chan<- int, q chan<- bool, num int) {
	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
}

//receive channels
func DiceResult(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case val := <-e:
			fmt.Println("from even--", val)
		case val := <-o:
			fmt.Println("from  odd--", val)
		case i, err := <-q:
			if err == true {
				fmt.Println("from quit--", i, err)
				return
			} else {
				fmt.Println("from quit--", err) //------channels q has false value
				return
			}
		}
	}
}
