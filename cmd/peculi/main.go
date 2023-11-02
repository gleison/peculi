package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gleison/peculi/web"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3333"
	}
	err := http.ListenAndServe(":"+port, mux)
	fmt.Printf("Error: %v\n", err)
}

func index(w http.ResponseWriter, r *http.Request) {
	index := web.GetHtmlFile("index.html")
	io.WriteString(w, string(index))
}
