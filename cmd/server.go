package main

import (
	"net/http"
)

func main() {
	r := NewRouter()
	http.ListenAndServe(":3000", r)
}
