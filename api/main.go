package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorld)
	log.Println("YO SERVER ON")
	http.ListenAndServe(":8080", nil)
}

func helloWorld(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello world\n")
}
