package testing

import (
	"fmt"
	"net/http"
	"testing"
)

func Test_net_http(t *testing.T) {
	fmt.Println("http://localhost:8000")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	http.ListenAndServe(":8000", nil)
}
