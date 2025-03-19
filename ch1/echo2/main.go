package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	// There are several ways to declare a string variable
	// s := "" // used only within function, not for package-level
	// var s string // rarely used except
	// var s = ""
	// var s string = ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	// use the blank identifier, whose name is _ (underscore)

	fmt.Println(s)

}
