package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// 1.8
	// const prefix string = "http://"

	for _, url := range os.Args[1:] {

		// 1.8
		// if !strings.HasPrefix(url, "http://") {
		// 	url = prefix + url
		// }

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: read %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%s", b)

		// 1.7
		// _, err = io.Copy(os.Stdout, resp.Body)
		// resp.Body.Close()
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v", url, err)
		// 	os.Exit(1)
		// }

	}
}
