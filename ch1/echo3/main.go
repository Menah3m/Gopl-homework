// Echo3 prints its commandline arguments.
// strings.Join()

package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

