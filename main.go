package main

import (
	"fmt"
	"html/template"
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

func loginController(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(rw, nil)
	} else {
		r.ParseForm()
		//logic part of login
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", indexController)
	http.HandleFunc("/login", loginController)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
