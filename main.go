package main

import (
	"fmt"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi there!!!")
}

func main() {
	http.HandleFunc("/", handleRoot)
	http.ListenAndServe(":3001", nil)
}
