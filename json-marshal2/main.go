package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	//Create the Json string
	var b = []byte(`
        {
        "id": 12423434, 
        "Name": "Fernando"
        }
    `)

	//Marshal the json to a map
	var f interface{}
	json.Unmarshal(b, &f)
	m := f.(map[string]interface{})

	//print the map
	fmt.Println(m)

	//unmarshal the map to json
	result, _ := json.Marshal(m)

	//print the json
	os.Stdout.Write(result)

}
