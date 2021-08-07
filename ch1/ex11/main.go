package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, args := range os.Args {
		s += sep + args
		sep = " "
	}
	fmt.Println(s)
}
