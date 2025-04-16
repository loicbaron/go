package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func storeHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse multipart form data (maxMemory specifies the max memory used before writing to disk)
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "Error parsing form data: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve the file from form data
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create a destination file
	dst, err := os.Create("./uploads/" + handler.Filename)
	if err != nil {
		http.Error(w, "Unable to create the file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the uploaded file's content to the destination file
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Error saving the file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully: %s\n", handler.Filename)
}

func main() {
	// Make sure uploads directory exists
	os.MkdirAll("./uploads", os.ModePerm)

	http.HandleFunc("/store", storeHandler)
	fmt.Println("Storage Server listening on :8181")
	http.ListenAndServe(":8181", nil)
}
