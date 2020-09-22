// read content by io.Copy() instead of ioutil.ReadAll()
package main

import (
	"io"
	"fmt"
	"os"
	"net/http"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
		}
		r, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "reading %s:%v\n", url, err)
		}
		fmt.Printf("%v", r)
	}
}