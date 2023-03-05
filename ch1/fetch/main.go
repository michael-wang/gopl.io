package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Exercise 1.7
var copy bool

// Exercise 1.8
var prefix bool

// Exercise 1.9
var status bool

func main() {
	flag.BoolVar(&copy, "copy", false, "Exercise 1.7 use io.Copy to copy response body to os.Stdout")
	flag.BoolVar(&prefix, "prefix", false, "Exercise 1.8 add prefix https:// if missing")
	flag.BoolVar(&status, "status", false, "Exercise 1.9 print the HTTP status code")
	flag.Parse()

	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, "--") {
			continue
		}
		if prefix && !strings.HasPrefix(url, "http") {
			url = "https://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		if copy {
			_, err := io.Copy(os.Stdout, resp.Body)
			resp.Body.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
				os.Exit(1)
			}
		} else {
			b, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
				os.Exit(1)
			}
			fmt.Printf("%s", b)
		}

		fmt.Println("fetch", url, "Done.")
		if status {
			fmt.Println("HTTP status code:", resp.Status)
		}
	}
}
