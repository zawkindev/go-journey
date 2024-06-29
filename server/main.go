package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

func showInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Current time: ", time.Now())
	fmt.Fprintln(w, "Url path: ", html.EscapeString(r.URL.Path))
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	http.HandleFunc("/", serveFile)
	http.HandleFunc("/info", showInfo)

	err := http.ListenAndServe(":8999", nil)
	if err != nil {
		log.Fatal(err)
	}
}
