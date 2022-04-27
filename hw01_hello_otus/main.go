package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	str := "Hello, OTUS!"

	fmt.Println(stringutil.Reverse(str))
}
