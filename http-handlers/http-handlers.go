package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (si String) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprintf(w, "StringHello! with %s", si)
}

func (su Struct) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	//fmt.Fprint(w, "StructHello! with ", su.Greeting, su.Punct, su.Who)
	fmt.Fprintf(w, "StructHello! with %v", su)
}

func main() {
	// your http.Handle calls here
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})

	log.Fatal(http.ListenAndServe("localhost:4000", nil))

	/*
	   var si String
	   err := http.ListenAndServe("localhost:4000", si)
	   if err != nil {
	     log.Fatal(err)
	   }
	*/
}
