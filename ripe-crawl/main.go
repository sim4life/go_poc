package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
	"text/scanner"

	"golang.org/x/net/publicsuffix"
	// third-party libraries
	"github.com/PuerkitoBio/goquery"
	iconv "github.com/djimenez/iconv-go"
)

// Custom type defined to make sure JSON output contains precision of 2 decimal points
type Number float32

func (n Number) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%.2f", n)), nil
}

// Basic struct to be used for channel sharing and JSON Marshaling
type FruitItem struct {
	Title       string `json:"title"`
	Size        string `json:"size"`
	UnitPrice   Number `json:"unit_price"`
	Description string `json:"description"`
	DetailsUri  string `json:"-"` //ignored on marshaling
}

// Fatal if there is an error
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

/**
 * This function parses a value string parameter and returns Number value
 * embedded within the string. It returns nil if it doesn't find any
 * Number value in the value string.
 * Example: "some4.56more" would return 4.56
 */
func extractFloat32(value string) Number {
	var sc scanner.Scanner
	var tok rune
	var valFloat64 float64
	var valFloat32 Number
	var err error
	var isFound bool

	if len(value) > 0 {
		sc.Init(strings.NewReader(value))
		sc.Mode = scanner.ScanFloats

		for tok != scanner.EOF {
			tok = sc.Scan()
			// fmt.Println("At position", sc.Pos(), ":", sc.TokenText())
			valFloat64, err = strconv.ParseFloat(sc.TokenText(), 64)
			if err == nil {
				isFound = true
				break
			}
		}
	}

	if isFound {
		valFloat32 = Number(valFloat64)
	}

	return valFloat32
}

/**
 * This function parses and returns the uri associated with the HTML anchor
 * <a href="http://www..."...> tag
 * This function assumes that 'href' attribute contains absolute url.
 * It returns "" empty string if it can't find href attribute from the
 * goquery.Selection parameter.
 */
func getUri(sel *goquery.Selection) string {
	if sel != nil {
		str, exists := sel.Attr("href")
		if exists {
			u, err := url.Parse(str)
			checkErr(err)
			return u.String()
		}
	}
	return ""
}

/**
 * This function scrapes the fruit item title and unit price from the downloaded
 * HTML document. First, it downloads the HTML doc from the given URI parameter.
 * Then, it scrapes the products' title, unit price and details URI info from the
 * downloaded HTML doc. Later on, it creates fruit items with these partial values,
 * and then it puts these fruit item objects into fruitInQueue channel.
 */
func fruitInitScrape(client *http.Client, uri string, fruitInQueue chan *FruitItem) {

	var price Number
	var addPrice Number
	var iter int

	// Load the URI
	res, err := client.Get(uri)
	checkErr(err)

	defer res.Body.Close()

	// Convert the "windows-1252" charset of the downloaded HTML to
	// utf-8 encoded HTML.
	utfBody, err := iconv.NewReader(res.Body, "windows-1252", "utf-8")
	checkErr(err)

	// use utfBody using goquery
	doc, err := goquery.NewDocumentFromReader(utfBody)
	checkErr(err)

	fmt.Println("about to find stuff\n")
	// Find required info within the document
	doc.Find("ul.productLister li").Each(func(i int, s *goquery.Selection) {
		product := s.Find(".productInner h3 a")
		title := strings.TrimSpace(product.Text())
		prodUri := getUri(product)

		priceStr := strings.TrimSpace(s.Find(".productInner p.pricePerUnit").Text())

		addProduct := s.Find(".crossSellInner h4.crossSellName a")
		addTitle := strings.TrimSpace(addProduct.Text())
		addPriceStr := strings.TrimSpace(s.Find(".crossSellInner p.pricePerUnit").Text())

		// Parsing float32 value
		price = extractFloat32(priceStr)

		iter++
		// Creating fruit item with partial values
		fruitItem := &FruitItem{Title: title, UnitPrice: price, Size: "0kb", Description: "", DetailsUri: prodUri}
		// Pretty-printing the progress of in channel processing
		fmt.Print("Found stuff")
		for i := 0; i < iter; i++ {
			fmt.Print("=")
		}
		fmt.Println(">\n")

		// Putting partially formed fruitItem on to fruitInQueue channel
		fruitInQueue <- fruitItem

		// These additional fruit items are the cross selling product items
		if len(addTitle) > 0 {
			addPrice = extractFloat32(addPriceStr)
			addProdUri := getUri(addProduct)

			iter++
			// Creating fruit item with partial values
			fruitItem := &FruitItem{Title: addTitle, UnitPrice: addPrice, Size: "0kb", Description: "", DetailsUri: addProdUri}
			// Pretty-printing the progress of in channel processing
			fmt.Print("Found stuff")
			for i := 0; i < iter; i++ {
				fmt.Print("=")
			}
			fmt.Println(">\n")

			// Putting partially formed fruitItem on to fruitInQueue channel
			fruitInQueue <- fruitItem
		}
	})

	fmt.Println("\nfinished finding stuff ... closing channel\n")
	// Closing the In channel as it is not needed
	close(fruitInQueue)

}

/**
 * This function retrieves the size (KB) of the downloaed HTML document (without assets)
 * and scrapes the description of the fruit item from downloaded HTML doc. First, it
 * consumes the fruit item struct from fruitInQueue channel. Then, it downloads the HTML
 * document to find and save size and description in the same fruit item struct. Lastly,
 * it puts this completely formed fruit item struct in the fruitOutQueue channel
 */
func fruitFinishScrape(client *http.Client, fruitInQueue, fruitOutQueue chan *FruitItem) {
	var iter int
	// Consuming fruitItem from fruitInQueue channel
	for fruitItem := range fruitInQueue {

		fmt.Println("about to find MORE stuff\n")
		iter++

		res, err := client.Get(fruitItem.DetailsUri)
		checkErr(err)

		defer res.Body.Close()
		// Reading to calculate the size of HTML document in kb
		body, err := ioutil.ReadAll(res.Body)
		checkErr(err)

		// Converting bytes into KB
		size := float32(len(body)) / float32(1024)

		// Restoring the io.ReadCloser to its original state to be re-read
		res.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		// Querying from the restored http.Response
		doc, err := goquery.NewDocumentFromResponse(res)
		checkErr(err)

		// Looking for fruit item description within newly downloaded HTML document
		doc.Find("#information").Each(func(i int, s *goquery.Selection) {
			desc := strings.TrimSpace(s.Find(".productText p").First().Text())

			fruitItem.Description = desc
			fruitItem.Size = strconv.FormatFloat(float64(size), 'f', 2, 32) + "kb"

			// Pretty-printing the progress of out channel processing
			fmt.Print("<")
			for i := 0; i < iter; i++ {
				fmt.Print("=")
			}
			fmt.Println("FOUND more stuff\n")

			// Putting completely formed fruitItem on to fruitOutQueue channel
			fruitOutQueue <- fruitItem

		})
	}
	fmt.Println("\nfinished finding MORE stuff ... closing channel\n")
	// Closing the Out channel as it is not needed
	close(fruitOutQueue)
}

/**
 * This function returns a json []byte from the fruit items. First, it
 * consumes fruit items from fruitOutQueue channel. Then it appends it to
 * fruitsList and finds the cumulative unit price of all the fruit items.
 * Later on, it builds the JSON object and returns it.
 */
func getFruitsJSON(fruitOutQueue chan *FruitItem) []byte {
	fruitsList := make([]*FruitItem, 0)
	totalPrice := Number(0)

	var iter int
	for fruitItem := range fruitOutQueue {
		iter++
		// Pretty-printing the progress of out channel processing
		fmt.Print("Thanks <")
		for i := 0; i < iter; i++ {
			fmt.Print("-")
		}
		fmt.Println("> stuff\n")

		fruitsList = append(fruitsList, fruitItem)
		totalPrice += fruitItem.UnitPrice
	}

	fmt.Println("Total Unit Price is: ", totalPrice)

	// A temporary struct is defined to generate desired JSON
	fruitsJSON, err := json.MarshalIndent(struct {
		Results    []*FruitItem `json:"results"`
		TotalPrice Number       `json:"total"`
	}{Results: fruitsList, TotalPrice: totalPrice},
		"", "   ")
	checkErr(err)

	return fruitsJSON

}

func main() {

	uri := "http://www.sainsburys.co.uk/webapp/wcs/stores/servlet" +
		"/CategoryDisplay?listView=true&orderBy=FAVOURITES_FIRST&" +
		"parent_category_rn=12518&top_category=12518&langId=44&" +
		"beginIndex=0&pageSize=20&catalogId=10137&searchTerm=&" +
		"categoryId=185749&listId=&storeId=10151&promotionId=#" +
		"langId=44&storeId=10151&catalogId=10137&categoryId=185749&" +
		"parent_category_rn=12518&top_category=12518&pageSize=20&" +
		"orderBy=FAVOURITES_FIRST&searchTerm=&beginIndex=0&" +
		"hideFilters=true"

	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, err := cookiejar.New(&options)
	checkErr(err)
	client := http.Client{Jar: jar}

	fmt.Println("the uri is: ", uri)

	fruitInQueue := make(chan *FruitItem, 2)
	fruitOutQueue := make(chan *FruitItem)
	go fruitInitScrape(&client, uri, fruitInQueue)
	go fruitFinishScrape(&client, fruitInQueue, fruitOutQueue)
	fruitsJSON := getFruitsJSON(fruitOutQueue)
	fmt.Println("JSON is: ", string(fruitsJSON))

}
