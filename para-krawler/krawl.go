// file: krawl.go
package main

import (
	"flag"
	"fmt"
	// "io/ioutil"
	"crypto/tls"
	"net/http"
	"net/url"
	"os"
	// "reflect"

	"github.com/jackdanger/collectlinks"
)

func retrieveQueue(client *http.Client, uri string, queue chan string) {
	fmt.Println("fetching: " + uri + " ...")
	var added = make(map[string]bool)
	resp, err := client.Get(uri)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	added[uri] = true
	// fmt.Println("uri visited is: ", uri)
	// fmt.Println("type of uri: ", reflect.TypeOf(uri))

	links := collectlinks.All(resp.Body)
	for _, link := range links {
		absolute := fixURL(link, uri)
		// fmt.Println("ab visited is: ", absolute)
		if absolute != "" && !added[absolute] {
			// fmt.Println("type of ab: ", reflect.TypeOf(absolute))
			added[absolute] = true
			go func() { queue <- absolute }()
		}
	}

}

func fixURL(href, base string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return ""
	}
	baseUrl, err := url.Parse(base)
	if err != nil {
		return ""
	}
	uri = baseUrl.ResolveReference(uri)
	return uri.String()
}

func main() {
	fmt.Println("I'm a gopher!")
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Please specify a start page to crawl")
		os.Exit(1)
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	client := http.Client{Transport: transport}
	/*
		resp, err := http.Get("http://6brand.com.com")
		body, err := ioutil.ReadAll(resp.Body)
	*/
	queue := make(chan string)
	go func() {
		queue <- args[0]
	}()

	for uri := range queue {
		retrieveQueue(&client, uri, queue)
	}

	//retrieve(args[0], &client)

}
