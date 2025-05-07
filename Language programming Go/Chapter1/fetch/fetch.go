package main

// go build fetch.go
// ./fetch http://gopl.io
// ./fetch http://bad.gopl.io

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			// fmt.Printf("\nStatus - %s\n", resp.Status)
			os.Exit(1)
		}
		bytes, err := io.Copy(os.Stdout, resp.Body)
		// b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("\n%d bytes readed, status - %s\n", bytes, resp.Status)
	}
}
