package main

import (
	"fmt"
	"time"
	"math/rand"
)


func main() {
	//var wg sync.WaitGroup
	d := PrintAllChannel(PrintName("hello"),PrintName("World"))

for i := 0; i < 100; i++ {
	fmt.Println(<-d)
}
//wg.Wait()

	}


	func PrintAllChannel(ch1, ch2 <-chan string)<-chan string{

		// var wg sync.WaitGroup
		// wg.Add(2)
		ch := make(chan string)
		go func()  {
			for {
				ch<-<-ch1
			}
//			wg.Done()
		}()
		go func()  {
			for {
				ch<-<-ch2
			}
//			wg.Done()
		}()
		return ch	
	}

	func PrintName(s string) <- chan string{
		cha := make(chan string)

		go func() {
			for i:= 0; ; i++ {
			cha <- fmt.Sprintf("%s...%d",s,i)
time.Sleep(time.Duration(rand.Intn(1e3))*time.Millisecond)
			
			}
		}()
		return cha
	}

