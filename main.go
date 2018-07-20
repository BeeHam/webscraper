package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	geturl := flag.String("url", "http://www.google.com", "get request destination")
	flag.Parse()

	resp, err := http.Get(*geturl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	respbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", respbody)
}
