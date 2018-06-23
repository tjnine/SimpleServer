package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func indexController(wr http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //parse arguments, you have to call this by yourself...
	fmt.Println(r.Form) // print info in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(wr, "A working net/http server using golangs' library") //send data to client side
}

func main() {
	http.HandleFunc("/", indexController)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
