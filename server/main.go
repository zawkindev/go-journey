package main

import (
	"log"
	"net/http"
)

type AfterMiddleware struct {
	handler http.Handler
}

func (a *AfterMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.handler.ServeHTTP(w, r)
	_, err := w.Write([]byte(" +++ Hello from middleware! +++ \n"))
	if err != nil {
		log.Fatal(err)
	}
}
func myHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(" *** Hello from myHandler! *** \n"))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	mid := &AfterMiddleware{http.HandlerFunc(myHandler)}
	println("Listening on port 8999")
	err := http.ListenAndServe(":8999", mid)
	if err != nil {
		log.Fatal(err)
	}
}
