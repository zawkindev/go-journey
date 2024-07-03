package main

import (
	"fmt"
	"log"
	"net/http"
)

type afterMiddleWare struct {
	handler http.Handler
}

func (a *afterMiddleWare) serveHttp(w http.ResponseWriter, r *http.Request) {
	a.handler.ServeHTTP(w, r)
	w.Write([]byte("Good Night dear"))
}

func middleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		fmt.Fprintln(w, "Executing middleware: 1")
	})
}

func middleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		fmt.Fprintln(w, "Executing middleware: 2")
	})
}

func final(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "final Middleware")
}

func main() {
	files := http.FileServer(http.Dir("./data/crucial"))

	mid := &afterMiddleWare{http.HandlerFunc(final)}

	http.Handle("/data/", http.StripPrefix("/data/", files))

	http.Handle("/", middleware1(middleware2(mid.handler)))

	err := http.ListenAndServe(":8999", nil)
	if err != nil {
		log.Fatal(err)
	}
}
