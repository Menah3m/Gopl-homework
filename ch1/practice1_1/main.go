// Practice1.1 prints its command name and arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, seq string
	fmt.Printf("command name: %v\n", os.Args[0])
	fmt.Println("command: ")
	for _, arg := range os.Args[1:] {
		s += seq + arg
		seq = " "
	}
	fmt.Println(s)
}
