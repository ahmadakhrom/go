package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {

	s := make(chan int)
	d := make(chan int)

	go Persons(s)
	go IDPersons(s, d)

	for v := range d {
		fmt.Println(v)
		fmt.Println("---------num of goroutines", runtime.NumGoroutine())
	}
	fmt.Println("---------num of processors", runtime.NumCPU())
}

func Persons(n chan int) {
	for i := 1; i <= 10; i++ {
		n <- i
	}
	close(n)
}

func IDPersons(s, d chan int) {

	var wg sync.WaitGroup

	const numOfGoRoutines = 100
	for i := 0; i < numOfGoRoutines; i++ {
		for sd := range s {
			wg.Add(1)

			go func(rs int) { //-----------------------------moving value of c1 to v2
				d <- timeConsumingWork(rs) //---------------"timeConsumingWork(rs)" is technic to (moving the value of c1 such as 123..n)
				wg.Done()                  //---------------onto random numbers that was resulted from function "timeConsumingWork"
			}(sd)
		}
	}
	wg.Wait()
	close(d)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
