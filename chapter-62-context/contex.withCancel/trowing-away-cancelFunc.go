package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println("printed as :\t\t", ctx)
	fmt.Printf("type of :\t\t %T\n", ctx)
	fmt.Println("have an error is :\t", ctx.Err())
	fmt.Println("---------------------------------")

	ctx, _ = context.WithCancel(ctx)
	fmt.Println("printed as :\t\t", ctx)
	fmt.Printf("type of :\t\t %T\n", ctx)
	fmt.Println("have an error is :\t", ctx.Err())

}
