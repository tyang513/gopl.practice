package golang

import (
	"fmt"
	"os"
)

func main() {
	var s = ""
	var sep = ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
