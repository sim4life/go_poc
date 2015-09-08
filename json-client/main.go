package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//	"os"
)

type Challenge struct {
	Token  string
	Values []int
}

func getJSON(url string, data interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
		log.Fatal("Error: %s", err)
	}
	defer resp.Body.Close()

	//fmt.Println("resp is:", resp)
	//fmt.Println("resp.Body is:", resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return err
	}
	fmt.Println("body is:", string(body))

	err = json.Unmarshal(body, data)
	if err != nil {
		//fmt.Println("error:", err)
		return err
	}

	return nil
}

func main() {

	var challenge Challenge
	url := "http://aerial-valor-93012.appspot.com/challenge"
	err := getJSON(url, &challenge)
	if err != nil {
		fmt.Printf("Error: %s", err)
		//os.Exit(1)
	}

	sum := 0
	for _, val := range challenge.Values {
		sum += val
	}
	fmt.Printf("%+v\n", challenge)
	fmt.Println("Sum is:", sum)

	url += "/%s/%d"
	url = fmt.Sprintf(url, challenge.Token, sum)
	var result map[string]interface{}
	getJSON(url, &result)
	fmt.Println("answer is:", result["answer"])

}
