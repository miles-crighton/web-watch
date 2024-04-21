package webwatch

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func traverseNode(n *html.Node, depth int) {
	indent := strings.Repeat("  ", depth)

	switch n.Type {
	case html.ElementNode:
		fmt.Printf("%s<%s>\n", indent, n.Data)
	case html.TextNode:
		text := strings.TrimSpace(n.Data)
		if text != "" {
			fmt.Printf("%s%s\n", indent, n.Data)
		}
	case html.CommentNode:
		fmt.Printf("%s<!-- %s -->\n", indent, n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverseNode(c, depth+1)
	}
}

func Fetch() {
	resp, err := http.Get("https://google.com")
	check(err)

	defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)
	check(err)
	// fmt.Printf("Response body: %v", string(body))

	err = os.MkdirAll("page-store", 0750)
	check(err)
	file, err := os.Create("page-store/google.com")
	check(err)
	defer file.Close()

	// file.Write(body)

	fmt.Println("Succesfully saved file")

	doc, err := html.Parse(io.Reader(resp.Body))
	check(err)
	traverseNode(doc, 0)
}
