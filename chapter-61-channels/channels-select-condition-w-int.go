package main

import "fmt"

func main() {

	fmt.Println("starting..")

	en := make(chan int)
	od := make(chan int)
	qu := make(chan int)
	num := 15

	//sending channels
	go Dice(en, od, qu, num)

	//receive channels
	DiceResult(en, od, qu)

	fmt.Println("done!")

}

//sending channels
func Dice(e, o, q chan<- int, num int) {
	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}

//receive channels
func DiceResult(e, o, q <-chan int) {
	for {
		select {
		case val := <-e:
			fmt.Println("from even--", val)
		case val := <-o:
			fmt.Println("from  odd--", val)
		case i, err := <-q:
			if err == true {
				fmt.Println("condition quit #1", i, err)
				return
			} else {
				fmt.Println("condition quit #2", i, err) //------channels q has false value
				return
			}
		}
	}
}
