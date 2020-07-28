package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Index)

	log.Fatal(http.ListenAndServe(":9001", nil))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Test Webpage")
}
