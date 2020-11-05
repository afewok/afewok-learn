package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("http://localhost:8000")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	http.ListenAndServe(":8000", nil)
}
