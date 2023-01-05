package main

import (
	"fmt"
	"net/http"
)

func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Kraken History")
	})

	http.ListenAndServe(":8080", nil)
}
