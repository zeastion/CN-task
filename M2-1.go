package main

import (
	"fmt"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Header)
}

func main() {
	http.HandleFunc("/", myHandler)
	http.ListenAndServe("127.0.0.1:80", nil)

}
