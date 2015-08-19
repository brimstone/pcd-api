package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hostname", handleHostname)

	fmt.Println("Starting http server on :8080")
	http.ListenAndServe("127.0.0.1:8080", nil)

}
