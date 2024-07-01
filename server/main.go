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

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Printf("%s %s %v\n", r.Method, r.URL.Path, time.Since(start))
	})
}

func main() {

	finalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})

	http.HandleFunc("/", finalHandler)
	http.HandleFunc("/info", showInfo)
	http.Handle("/log", logger(finalHandler))

	files := http.FileServer(http.Dir("C:\\Users\\shahruz\\webserver\\www"))
	http.Handle("/files/", http.StripPrefix("/files/", files))
	err := http.ListenAndServe(":8999", nil)
	if err != nil {
		log.Fatal(err)
	}
}
