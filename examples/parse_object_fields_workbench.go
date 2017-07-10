package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/net/html"
)

// Reference using:
// https://stackoverflow.com/questions/30109061/golang-parse-html-extract-all-content-with-body-body-tags

func getBody(doc *html.Node) (*html.Node, error) {
	var b *html.Node
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "body" {
			b = n
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	if b != nil {
		return b, nil
	}
	return nil, errors.New("Missing <body> in the node tree")
}

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}

func buildObjects() {
	filename := "_parse_object_fields_workbench.html"
	r, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	if 1 == 0 {
		doc, err := html.Parse(r)
		if err != nil {
			log.Fatal(err)
			fmt.Println("ERR")
		}
		fmt.Printf("%v\n", doc)
		fmt.Println("SUC")

		bn, err := getBody(doc)
		if err != nil {
			return
		}
		body := renderNode(bn)
		fmt.Println(body)
	}

	if 1 == 0 {
		z := html.NewTokenizer(r)

		for {
			tt := z.Next()
			if tt == html.ErrorToken {
				// ...
				fmt.Println("ERR")
			}
			// Process the current token.
		}
	}
}

func main() {
	buildObjects()
	fmt.Println("DONE")
}
