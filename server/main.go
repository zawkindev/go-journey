package main

import (
	"fmt"
	"net/http"
)

func middleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Executing middleware: 1")
		next.ServeHTTP(w, r)
		fmt.Fprintln(w, "Executing middleware: 1 AGAIN")
	})
}

func middleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Executing middleware: 2")
		next.ServeHTTP(w, r)
		fmt.Fprintln(w, "Executing middleware: 2 AGAIN")
	})
}

func final(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "final Middleware")
	fmt.Fprintln(w, "Done!")
}

func main() {
	http.Handle("/", middleware1(middleware2(http.HandlerFunc(final))))
	http.ListenAndServe(":8999", nil)
}
