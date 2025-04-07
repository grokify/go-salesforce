package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log/slog"
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
	return nil, errors.New("missing <body> in the node tree")
}

func renderNode(n *html.Node) (string, error) {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	if err := html.Render(w, n); err != nil {
		return "", err
	} else {
		return buf.String(), nil
	}
}

func buildObjects() error {
	filename := "_parse_object_fields_workbench.html"
	r, err := os.Open(filename)
	if err != nil {
		return err
	}

	if 1 == 0 {
		doc, err := html.Parse(r)
		if err != nil {
			return err
		}
		slog.Info("success_on_parse", "doc", fmt.Sprintf("%v", doc))

		bn, err := getBody(doc)
		if err != nil {
			return err
		}
		body, err := renderNode(bn)
		if err != nil {
			return err
		}
		slog.Info("success_on_render", "body", body)
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
	return nil
}

func main() {
	err := buildObjects()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	fmt.Println("DONE")
	os.Exit(0)
}
