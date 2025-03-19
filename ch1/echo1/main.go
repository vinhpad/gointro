package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	// The var declaration type of string.
	// A variable can be initialized
	// as part of declaration (optional).
	// The variable is implicitly initialized
	// to the zero value for its type ( 0 for int , "" for string )

	for i := 0; i < len(os.Args); i++ {
		println(os.Args[i])
		s += sep + os.Args[i]
		sep = " "
	}

	// i += 1 like i = i + 1, i++, ++i but j = i++ is illegal

	fmt.Println(s)

	// Loop illustrate here:
	// for initialization; condition; post {
	//		// zero or more statements
	// }

	// A tradition "while" loop:
	// for condition {
	//		// ...
	// }
}
