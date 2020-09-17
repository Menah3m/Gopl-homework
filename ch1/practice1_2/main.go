// practice1_2 prints its command's index and value.
package main

import (
	"fmt"
	"os"
)

func main() {

	for ind, val := range os.Args[1:] {
		fmt.Printf("index: %v, value: %v \n", ind, val)
	}
}
