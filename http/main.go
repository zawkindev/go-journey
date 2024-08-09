package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://github.com/zawkindev/go-journey")
	if err != nil {
		log.Fatal(err)
	}
  defer resp.Body.Close()

  fmt.Println("status: ", resp.Status)

  scanner := bufio.NewScanner(resp.Body)
  for scanner.Scan(){
    fmt.Println(scanner.Text())
  }
}
