package functions

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

//FindLinks is a func to find link on the web-pages
func FindLinks(url string) ([]string, error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		res.Body.Close()

		return nil, fmt.Errorf("getting: %s: %s", url, res)
	}

	doc, err := html.Parse(res.Body)
	res.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	return VisitPage(nil, doc), nil
}
