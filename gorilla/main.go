package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func pageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]
	log.Printf("Product ID: %v\n", productID)

	filename := productID + ".html"

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Println("404: file not found")
		http.Redirect(w, r, "/404", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, filename)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	http.ServeFile(w, r, "404.html")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/product/{id:[0-9]+}", pageHandler)
	router.HandleFunc("/404", notFoundHandler)
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	http.Handle("/", router)
	http.ListenAndServe(":8999", nil)
}
