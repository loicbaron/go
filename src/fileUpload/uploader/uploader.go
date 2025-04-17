package main

import (
	"fileUpload/logger"
	"fmt"
	"io"
	"net/http"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info.Printf("received request %s %s %s", r.Method, r.URL.Path, r.RemoteAddr)
	// Only allow POST requests
	if r.Method != http.MethodPost {
		errMsg := fmt.Sprintf("Method %s not allowed", r.Method)
		logger.Error.Print(errMsg)
		http.Error(w, errMsg, http.StatusMethodNotAllowed)
		return
	}

	// Prepare a new request to the target server
	targetURL := "http://localhost:8181/store"

	// Create a new POST request with the same body
	req, err := http.NewRequest(http.MethodPost, targetURL, r.Body)
	if err != nil {
		errMsg := "Failed to create forward request: "+err.Error()
		logger.Error.Print(errMsg)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}
	// Important: copy the headers (especially Content-Type for multipart/form-data)
	req.Header = r.Header

	// Use default HTTP client to send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errMsg := "Error forwarding request: "+err.Error()
		logger.Error.Print(errMsg)
		http.Error(w, errMsg, http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	// Copy response headers and body from target server to client
	for key, values := range resp.Header {
		for _, v := range values {
			w.Header().Add(key, v)
		}
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func main() {
	logger.Init()

	http.HandleFunc("/upload", uploadHandler)
	logger.Info.Println("Upload Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
