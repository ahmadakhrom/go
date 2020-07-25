package main

import 	"fmt"

func main() {
	c := make(chan int)
//simple way
go func(){
	c<-44	
}()

v, ok := <-c
fmt.Println(v, ok)

close(c)

v, ok = <-c
fmt.Println(v, ok)   // 0 false coz channel was closed

 }