package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

func main() {
	counts := make(map[string]int)

	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)

	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	// [value, key]
	for line, n := range counts {
		fmt.Printf("%d\t%s\t%d\n", n, line, counts[line])
	}
}

// %d 			decimal integer
// %x, %o, %b 	integer in hexadecimal, octal, binary
// %f, %g, %e	floating-point number
// %t			boolean: true or false
// %c			rune (unicode)
// %s			string
// %q			quoted string "abc" or rune 'c'
// %v			any value
// %T 			type of any value
// %%			literal percent sign
