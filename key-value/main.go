package main

import (
	"context"
	"fmt"
)


func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "hello", "world")
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	world := ctx.Value("hello")
	fmt.Println(world)
}