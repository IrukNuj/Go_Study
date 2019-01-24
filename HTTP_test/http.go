package main

import (
	"fmt"
	"net/http"
)


func handler(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	fmt.Fprintf(w, "Hello!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
