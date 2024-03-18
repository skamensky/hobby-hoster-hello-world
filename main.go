package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	fmt.Println("Started hello world server!")
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":80", nil)
}
