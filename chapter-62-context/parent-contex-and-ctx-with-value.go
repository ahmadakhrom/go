package main

import (
	"context"	
	"fmt"
)


func main()  {
	
	ctx := context.WithValue(context.Background(),"color","yellow")
	foo(ctx, "color")

	ctx = context.WithValue(context.Background(),"color","white")
	foo(ctx, "color")

	foo(ctx, "animal")
}

func foo(ctx context.Context, s string){
	if v := ctx.Value(s); v != nil {
		fmt.Println("found value :",v)
		return
	}
	fmt.Println("not found :",s)	
}