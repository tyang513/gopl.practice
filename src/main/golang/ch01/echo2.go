package main

import (
	"fmt"
	"os"
)

func main() {
	var s = ""
	var sep = ""
	for a, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		fmt.Println(a)
	}
	fmt.Println(s)
}
