package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
    port := os.Getenv("PORT")
    if port == "" {
        port = "3333"
    }
	err := http.ListenAndServe(":" + port, mux)
	fmt.Printf("Error: %v\n", err)
}

func index(w http.ResponseWriter, r *http.Request) {
	index, _ := os.ReadFile("web/index.html")
	io.WriteString(w, string(index))
}
