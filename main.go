package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	// retrieve url to target
	geturl := flag.String("url", "http://www.google.com/robots.txt", "get request destination")
	flag.Parse()
	// send GET request and assign response
	resp, err := http.Get(*geturl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// tokenize response body
	tokBody := html.NewTokenizer(resp.Body)
	// loop through all tokens
	for {
		tokType := tokBody.Next()

		switch {
		case tokType == html.ErrorToken:
			// if Errortoken received, we've come to EOF
			return
		case tokType == html.StartTagToken:
			if tok := tokBody.Token(); tok.Data == "a" {
				// if we've found an <a> tag, iterate through all of its attributes
				// and locate the href
				for _, v := range tok.Attr {
					if v.Key == "href" {
						fmt.Println("Found href: ", v.Val)
						break
					}
				}
			}
		}
	}

}
