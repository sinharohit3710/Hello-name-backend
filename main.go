package main

import (
	"fmt"
	"net/http"
	"os"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s", os.Getenv("NAME"))
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
}

func main() {
	setupRoutes()
	http.ListenAndServe(":3000", nil)
}
