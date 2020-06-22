package main

import "fmt"

func main(){
	fmt.Println("Hello world..")
	
	foo()
	
	fmt.Println("sometin' more")

	x, y:= 5, 100
	for i:=0; i<y;i++ {			//y adalah nilai maksimal perulangan
		if i%x == 0{			//mencari angka dari kelipatan value variabel x
			fmt.Println(i)
		} 
	}

	bar()

}

func foo()  {
fmt.Println("i'm in foo function")	
}

func bar()  {
	fmt.Println("program exited")
	
}