package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	s, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Printf("Error reading stdin: %v\n", err)
		os.Exit(1)
	}

	steps := string(s)
	floor := 0

	for i, c := range steps {
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
		}
		if floor == -1 {
			fmt.Println(i + 1)
			return
		}
	}

}
