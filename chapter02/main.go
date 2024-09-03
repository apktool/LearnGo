package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, URL.Path = %q\n", r.URL.Path)
	})
	res := http.ListenAndServe(":8080", nil)
	log.Fatal(res)
}
