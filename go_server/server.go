package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", IndexServer)
	http.ListenAndServe(":8000", nil)
}

func IndexServer(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello from go!")
}