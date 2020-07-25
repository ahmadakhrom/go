package main

import (
	"context"
	"fmt"
)


func main() {
	ctx := context.Background()

	fmt.Println("whats in context.background\t\t: ",ctx)
	fmt.Println("whats in err on context.background\t: ",ctx.Err())
	fmt.Printf("whats is type context.background\t: %T",ctx)
}