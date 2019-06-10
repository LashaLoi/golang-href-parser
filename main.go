package main

import (
	"fmt"
	"os"

	"./functions"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err := functions.FindLinks(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "parse: %v\n", err)
		}

		for _, links := range links {
			fmt.Println(links)
		}
	}
}
