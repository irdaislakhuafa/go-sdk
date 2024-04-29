package main

import (
	"context"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
	"github.com/irdaislakhuafa/go-sdk/log"
)

func main() {
	l := log.Init(log.Config{Level: "debug"})
	err := errors.NewWithCode(codes.CodeBadRequest, "haha")
	l.Debug(context.Background(), err)
}
