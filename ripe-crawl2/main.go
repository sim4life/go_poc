package main

import (
	// "crypto/tls"
	"fmt"
	// "io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	// "net/url"
	// "strconv"
	"strings"
	// "text/scanner"

	"golang.org/x/net/publicsuffix"
	// third-party libraries
	"github.com/PuerkitoBio/goquery"
)

// fatal if there is an error
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func retrieveSizeDesc(uri string) {
	fmt.Printf("uri is: %s\n", uri)

	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, err := cookiejar.New(&options)
	checkErr(err)
	client := http.Client{Jar: jar}

	// Load the URL
	// res, err := http.Get(uri)
	res, err := client.Get(uri)
	checkErr(err)

	defer res.Body.Close()

	/*
		body, err := ioutil.ReadAll(res.Body)
		checkErr(err)
		fmt.Println("before body is:\n", string(body))
	*/
	// use utfBody using goquery
	// doc, err := goquery.NewDocumentFromReader(utfBody)
	// doc, err := goquery.NewDocument(uri)
	doc, err := goquery.NewDocumentFromResponse(res)

	// doc, err := goquery.NewDocumentFromResponse(res)
	checkErr(err)
	fmt.Println("doc MORE is:\n", doc)
	// fmt.Println("doc.Contents is:\n", doc.Contents())
	fmt.Println("about to find MORE stuff\n")
	doc.Find("#information").Each(func(i int, s *goquery.Selection) {
		desc := strings.TrimSpace(s.Find(".productText p").First().Text())

		fmt.Println("prod desc is: ", desc)

		// fruitItem.description = desc

		// fruitOutQueue <- fruitItem

	})
}

func main() {
	uri := "http://www.sainsburys.co.uk/shop/gb/groceries/ripe---ready/sainsburys-conference-pears--ripe---ready-x4-%28minimum%29"
	retrieveSizeDesc(uri)
}
