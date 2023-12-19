package main

import (
	"context"
	"reflect"
	"time"

	"github.com/irdaislakhuafa/go-sdk/log"
)

type Person struct {
	ID        int64
	Code      string
	Name      string
	Age       int64
	BirthDate time.Time
	IsAlive   bool
}

func main() {
	ctx := context.Background()
	l := log.Init(log.Config{Level: "debug"})

	a := Person{ID: 1, Code: "xxx", Name: "xxx", Age: 21, IsAlive: true}
	b := Person{ID: 1, Code: "xxx", Name: "xxx", Age: 21, IsAlive: true}

	if reflect.DeepEqual(a, b) {
		l.Info(ctx, "is equals")
	} else {
		l.Debug(ctx, "is not equals")
	}
}
