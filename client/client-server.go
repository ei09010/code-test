package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	fmt.Println("client on localhost:7777")
	log.Fatal(http.ListenAndServe(":7777", nil))
}
