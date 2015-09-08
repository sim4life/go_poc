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

type FruitItem struct {
	Title       string  `json:"title"`
	UnitPrice   float32 `json:"unit_price"`
	Size        float32 `json:"size"`
	Description string  `json:"description"`
	Uri         string  `json:"-"` //ignored on marshalling
}

// fatal if there is an error
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func extractFloat(value string) float32 {
	var sc scanner.Scanner
	var tok rune
	var valFloat64 float64
	var valFloat32 float32
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
			// checkErr(err)
		}
	}

	if isFound {
		valFloat32 = float32(valFloat64)
	}

	return valFloat32
}

func getUri(sel *goquery.Selection) string {
	if sel != nil {
		str, exists := sel.Attr("href")
		// fmt.Println("str is: ", str)
		if exists {
			u, err := url.Parse(str)
			// fmt.Println("u is: ", u)
			checkErr(err)
			return u.String()
		}
	}

	return ""
}

func productScrape(client *http.Client, uri string, fruitInQueue chan *FruitItem) {

	// Load the URL
	// res, err := http.Get(uri)
	res, err := client.Get(uri)
	checkErr(err)

	defer res.Body.Close()

	// body, err := ioutil.ReadAll(res.Body)
	// checkErr(err)
	// fmt.Println("before body is:\n", string(body))

	// Convert the designated charset HTML to utf-8 encoded HTML.
	// `charset` being one of the charsets known by the iconv package.
	utfBody, err := iconv.NewReader(res.Body, "windows-1252", "utf-8")
	checkErr(err)

	// use utfBody using goquery
	doc, err := goquery.NewDocumentFromReader(utfBody)
	// doc, err := goquery.NewDocument(uri)
	// doc, err := goquery.NewDocumentFromResponse(res)
	checkErr(err)
	var price float32
	var addPrice float32
	var iter int

	fmt.Println("doc is:\n", doc)
	// fmt.Println("doc.Contents is:\n", doc.Contents())
	fmt.Println("about to find stuff\n")
	doc.Find("ul.productLister li").Each(func(i int, s *goquery.Selection) {
		product := s.Find(".productInner h3 a")
		title := strings.TrimSpace(product.Text())
		prodUri := getUri(product)

		priceStr := strings.TrimSpace(s.Find(".productInner p.pricePerUnit").Text())

		addProduct := s.Find(".crossSellInner h4.crossSellName a")
		addTitle := strings.TrimSpace(addProduct.Text())
		addPriceStr := strings.TrimSpace(s.Find(".crossSellInner p.pricePerUnit").Text())

		price = extractFloat(priceStr)

		// fmt.Printf("Title %d: %s - %.2f\n", i, title, price)
		// fmt.Printf("URI: %s\n", prodUri)
		iter++
		fruitItem := &FruitItem{Title: title, UnitPrice: price, Size: 0, Description: "", Uri: prodUri}
		// fmt.Println("Found stuff=============>\n")
		fmt.Print("Found stuff")
		for i := 0; i < iter; i++ {
			fmt.Print("=")
		}
		fmt.Println(">\n")

		// go func() {
		fruitInQueue <- fruitItem
		// }()
		if len(addTitle) > 0 {
			addPrice = extractFloat(addPriceStr)
			addProdUri := getUri(addProduct)
			// fmt.Printf("Additional Title %d: %s - %.2f\n", i, addTitle, addPrice)
			// fmt.Printf("Additional URI: %s\n", addProdUri)

			iter++
			fruitItem := &FruitItem{Title: addTitle, UnitPrice: addPrice, Size: 0, Description: "", Uri: addProdUri}
			// fmt.Println("Found stuff============>\n")
			fmt.Print("Found stuff")
			for i := 0; i < iter; i++ {
				fmt.Print("=")
			}
			fmt.Println(">\n")

			// go func() {
			fruitInQueue <- fruitItem
			// }()

		}
	})

	close(fruitInQueue)
	fmt.Println("\nfinished finding stuff\n")

}

func retrieveSizeDesc(client *http.Client, fruitInQueue, fruitOutQueue chan *FruitItem) {
	var iter int
	for fruitItem := range fruitInQueue {
		// fmt.Printf("In is: %+v\n", fruitItem)
		fmt.Println("about to find MORE stuff\n")
		iter++
		res, err := client.Get(fruitItem.Uri)
		checkErr(err)

		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		checkErr(err)
		// fmt.Println("before body is:\n", len([]byte(string(body))))
		// Converting bytes into kb
		size := float32(len(body)) / float32(1024)
		// Restore the io.ReadCloser to its original state to be re-read
		res.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		// fmt.Printf("len(body) is: %.2f\n", size)
		// doc, err := goquery.NewDocument(fruitItem.uri)
		doc, err := goquery.NewDocumentFromResponse(res)
		checkErr(err)
		// fmt.Println("doc MORE is:\n", doc)

		// fmt.Println("HTML.size is:\n", res.Body)

		// fmt.Println("doc.Contents is:\n", doc.Contents())

		doc.Find("#information").Each(func(i int, s *goquery.Selection) {
			desc := strings.TrimSpace(s.Find(".productText p").First().Text())

			// fmt.Println("prod desc is: ", desc)

			fruitItem.Description = desc
			fruitItem.Size = size

			// fmt.Printf("In is: %+v\n", fruitItem)
			// fmt.Printf("len(OutQueue) is: %d\n", len(fruitOutQueue))
			fmt.Print("<")
			// fmt.Println("<============FOUND stuff\n")
			for i := 0; i < iter; i++ {
				fmt.Print("=")
			}
			fmt.Println("FOUND more stuff\n")
			// go func(fruitItem *FruitItem) {
			fruitOutQueue <- fruitItem
			// }(fruitItem)

		})
	}
	fmt.Println("\nfinished finding MORE stuff\n")
	close(fruitOutQueue)
}

func buildJSON(fruitOutQueue chan *FruitItem) {
	fruitList := make([]*FruitItem, 0)
	totalPrice := float32(0)
	var iter int
	for fruitItem := range fruitOutQueue {
		iter++
		// fmt.Println("Thanks <---> stuff\n")
		fmt.Print("Thanks <")
		for i := 0; i < iter; i++ {
			fmt.Print("-")
		}
		fmt.Println("> stuff\n")

		// fmt.Printf("OutItem is: %+v\n", fruitItem)
		fruitList = append(fruitList, fruitItem)
		totalPrice += fruitItem.UnitPrice
	}
	/*
		for i, fruitItem := range fruitList {
			fmt.Printf("ListItem[%d] is: %+v\n", i, fruitItem)
			/*
				fruitJSON, err := json.Marshal(*fruitItem)
				checkErr(err)
				fmt.Println("JSON is: ", string(fruitJSON))

		}
	*/
	type FruitsTotal struct {
		Results    []*FruitItem `json:"results"`
		TotalPrice float32      `json:"total"`
	}
	fmt.Println("Total Unit Price is: ", totalPrice)
	// fruitsJSON, err := json.Marshal(fruitList)
	fruitsTotal := &FruitsTotal{Results: fruitList, TotalPrice: totalPrice}
	fruitsJSON, err := json.MarshalIndent(fruitsTotal, "", "    ")
	checkErr(err)
	// finalJSON := []byte(`{"results":`)
	// finalJSON = append(finalJSON, fruitsJSON...)
	// endJSON := []byte(`,"total":`)
	fmt.Println("JSON is: ", string(fruitsJSON))

}

func main() {
	// ExampleScrape()

	uri := "http://www.sainsburys.co.uk/webapp/wcs/stores/servlet" +
		"/CategoryDisplay?listView=true&orderBy=FAVOURITES_FIRST&" +
		"parent_category_rn=12518&top_category=12518&langId=44&" +
		"beginIndex=0&pageSize=20&catalogId=10137&searchTerm=&" +
		"categoryId=185749&listId=&storeId=10151&promotionId=#" +
		"langId=44&storeId=10151&catalogId=10137&categoryId=185749&" +
		"parent_category_rn=12518&top_category=12518&pageSize=20&" +
		"orderBy=FAVOURITES_FIRST&searchTerm=&beginIndex=0&" +
		"hideFilters=true"

		//uri2 := "http://www.sainsburys.co.uk/webapp/wcs/stores/servlet/CategoryDisplay?listView=true&orderBy=FAVOURITES_FIRST&parent_category_rn=12518&top_category=12518&langId=44&beginIndex=0&pageSize=20&catalogId=10137&searchTerm=&categoryId=185749&listId=&storeId=10151&promotionId=#langId=44&storeId=10151&catalogId=10137&categoryId=185749&parent_category_rn=12518&top_category=12518&pageSize=20&orderBy=FAVOURITES_FIRST&searchTerm=&beginIndex=0&hideFilters=true"

	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, err := cookiejar.New(&options)
	checkErr(err)
	client := http.Client{Jar: jar}

	fmt.Println("the uri is: ", uri)
	//fmt.Println("the uri2 is:", uri2)
	fruitInQueue := make(chan *FruitItem, 2)
	fruitOutQueue := make(chan *FruitItem)
	go productScrape(&client, uri, fruitInQueue)
	go retrieveSizeDesc(&client, fruitInQueue, fruitOutQueue)
	buildJSON(fruitOutQueue)

}
