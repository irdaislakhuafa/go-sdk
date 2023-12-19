package main

import (
	"fmt"

	"github.com/irdaislakhuafa/go-sdk/cryptography"
)

func main() {
	s, _ := cryptography.NewSHA256([]byte("password")).WithKey([]byte("key")).Build()
	fmt.Printf("s: %v\n", s)
}
