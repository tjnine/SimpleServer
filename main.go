package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func indexController(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //parse arguments, you have to call this by yourself...
	fmt.Println(r.Form) // print info in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(rw, "Welcome to the IndexController of the application.") //send data to client side
}

func login(rw http.ResponseWriter, r *http.Request) {
	// missing a bunch of code!
	r.ParseForm()
}

func main() {
	http.HandleFunc("/", indexController)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
