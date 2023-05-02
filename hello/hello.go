package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	fmt.Println(stringutil.ToUpper("Hello"))
	x := 42
	fmt.Println(x)
	x = 99
	y := 100 + 24
	fmt.Println(y)
}
