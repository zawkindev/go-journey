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
		log.Printf("404: file not found")
		filename = "404.html"
	}

	http.ServeFile(w, r, filename)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/product/{id:[0-9]+}", pageHandler)
	http.Handle("/", router)
	http.ListenAndServe(":8999", nil)
}
