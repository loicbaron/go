package main

import (
	"fileUpload/logger"
	"fmt"
	"io"
	"net/http"
	"os"
)

func storeHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info.Printf("received request %s %s %s", r.Method, r.URL.Path, r.RemoteAddr)
	// Only allow POST requests
	if r.Method != http.MethodPost {
		errMsg := fmt.Sprintf("Method %s not allowed", r.Method)
		logger.Error.Print(errMsg)
		http.Error(w, errMsg, http.StatusMethodNotAllowed)
		return
	}

	// Parse multipart form data (maxMemory specifies the max memory used before writing to disk)
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		errMsg := "Error parsing form data: " + err.Error()
		logger.Error.Print(errMsg)
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	// Retrieve the file from form data
	file, handler, err := r.FormFile("file")
	if err != nil {
		errMsg := "Error retrieving the file: " + err.Error()
		logger.Error.Print(errMsg)
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create a destination file
	dst, err := os.Create("./uploads/" + handler.Filename)
	if err != nil {
		errMsg := "Unable to create the file: " + err.Error()
		logger.Error.Print(errMsg)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the uploaded file's content to the destination file
	_, err = io.Copy(dst, file)
	if err != nil {
		errMsg := "Error saving the file: " + err.Error()
		logger.Error.Print(errMsg)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}
	successMsg := fmt.Sprintf("File stored successfully: %s\n", handler.Filename)
	logger.Info.Println(successMsg)
	fmt.Fprintf(w, "%s", successMsg)
}

func main() {
	logger.Init()
	// Make sure uploads directory exists
	os.MkdirAll("./uploads", os.ModePerm)

	http.HandleFunc("/store", storeHandler)
	logger.Info.Println("Storage Server listening on :8181")
	http.ListenAndServe(":8181", nil)
}
