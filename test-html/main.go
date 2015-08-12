package main

import (
	"fmt"
	"net/http"
)

func HelloWorld(res http.ResponseWriter, req *http.Request) {
	fmt.Println("key is " + req.FormValue("key"))
	fmt.Fprint(res, "Hello World "+req.FormValue("key"))
}

func main() {
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":3000", nil)
}
