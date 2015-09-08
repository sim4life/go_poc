package main

import (
	"encoding/json"
	"fmt"
	"log"
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

func buildJSON(fruitItem *FruitItem) {
	/*
	  for _, fruitItem := range fruitList {
	    // fmt.Printf("ListItem[%d] is: %+v\n", i, fruitItem)
	    fruitJSON, err := json.Marshal(*fruitItem)
	    checkErr(err)
	    fmt.Println("JSON is: ", string(fruitJSON))
	  }
	*/
	// fmt.Println("Total Unit Price is: ", totalPrice)
	fruitsJSON, err := json.Marshal(fruitItem)
	checkErr(err)
	fmt.Println("JSON is: ", string(fruitsJSON))

}

func main() {
	fruitItem1 := &FruitItem{Title: "Pear", UnitPrice: 3.5, Size: 33.8, Description: "Sweet Pears", Uri: "http://www.testing.com"}
	// fruitItem2 := &FruitItem{title: "Apple", unitPrice: 2.0, size: 35.1, description: "Golden Apples", uri: "http://www.testing2.com"}

	buildJSON(fruitItem1)

}
