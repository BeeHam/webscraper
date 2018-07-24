package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	geturl := flag.String("url", "http://www.google.com/robots.txt", "get request destination")
	flag.Parse()

	resp, err := http.Get(*geturl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// respbody, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// fmt.Printf("%s", respbody)

	z := html.NewTokenizer(resp.Body)
	for {
		tokType := z.Next()

		switch {
		case tokType == html.ErrorToken:
			return
		case tokType == html.StartTagToken:
			if tok := z.Token(); tok.Data == "a" {
				fmt.Println("Link found.")
			}
		default:
			fmt.Println("No link found.")
		}
	}

}
