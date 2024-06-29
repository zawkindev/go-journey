package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.ListenAndServe(":8999", http.FileServer(http.Dir("/var/www")))
	fmt.Print("salom")
}
