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

	currentUser, err := user.Current()
	if err != nil {
		http.Error(w, "Failed to get current user's information", http.StatusInternalServerError)
		return
	}

	dataDir := filepath.Join(currentUser.HomeDir, "data")

	// Create the directory if it doesn't exist
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		if err := os.MkdirAll(dataDir, 0700); err != nil {
			http.Error(w, "Failed to create data directory", http.StatusInternalServerError)
			return
		}
	}

	filePath := filepath.Join(dataDir, req.URL.Path[len("/fs/"):])
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
