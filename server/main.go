package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHello(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(rw, r.Form)
	fmt.Fprintln(rw, r.Form)
	fmt.Fprintln(rw, "path: ", r.URL.Path)
	fmt.Fprintln(rw, "scheme: ", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Fprintln(rw, "key:", k)
		fmt.Fprintln(rw, "val:", strings.Join(v, ", "))
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
