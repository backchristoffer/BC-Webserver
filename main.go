package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func bcw(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "bc webserver\n")
}

func fs(w http.ResponseWriter, req *http.Request) {
	directory := "."
	filePath := filepath.Join(directory, req.URL.Path[len("/fs/"):])

	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		http.Error(w, fmt.Sprintf("File '%s' not found", req.URL.Path), http.StatusNotFound)
		return
	}

	http.ServeFile(w, req, filePath)
}

func main() {
	http.HandleFunc("/", bcw)
	http.HandleFunc("/fs/", fs)
	port := 8080
	fmt.Printf("bc webserver running on port %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
