package main

import (
	"context"
	"fmt"

	"github.com/irdaislakhuafa/go-sdk/concurrency"
)

func main() {
	c := concurrency.NewConcurrency().WithMaxWorker(2)
	for i := 1; i <= 10; i++ {
		c.AddFunc(func(ctx context.Context, c concurrency.Interface) {
			fmt.Println(i)
		})
	}

	if err := c.Do(context.Background()); err != nil {
		panic(err)
	}

}
