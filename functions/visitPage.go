package functions

import "golang.org/x/net/html"

//VisitPage is a func for visit web apges
func VisitPage(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = VisitPage(links, c)
	}

	return links
}
