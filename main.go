package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
 	for _, url = range os.Args([1:]) {
		links, err = findLinks(url)	 
	
		if err != nil {
			fmt.Fprintf(os.Stderr, "parse: %v\n", err);
		}
	}
}