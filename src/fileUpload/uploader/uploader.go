package main

import (
	"fileUpload/logger"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
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
		errMsg := "Failed to create forward request: " + err.Error()
		logger.Error.Print(errMsg)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}
	// Important: copy the headers (especially Content-Type for multipart/form-data)
	req.Header = r.Header

	// Initialize the HTTP client
	client := &http.Client{}
	maxRetries := 5
	backoffFactor := time.Second
	var resp *http.Response

	// Retry loop with exponential backoff
	for attempt := 0; attempt < maxRetries; attempt++ {
		resp, err = client.Do(req)
		if err == nil && resp.StatusCode < 500 {
			// If request is successful or a non-retryable error occurs (i.e., 2xx or 4xx),
			// break out of the retry loop
			break
		}

		// If there's an error or a 5xx response, we retry
		if err != nil {
			errMsg := fmt.Sprintf("Error forwarding request: %s", err.Error())
			logger.Error.Print(errMsg)
		} else {
			errMsg := fmt.Sprintf("Received 5xx status %d, retrying...", resp.StatusCode)
			logger.Error.Print(errMsg)
		}

		// If we've exhausted the retry limit, we return a failure
		if attempt == maxRetries-1 {
			errMsg := fmt.Sprintf("Max retries reached for request to %s", targetURL)
			logger.Error.Print(errMsg)
			http.Error(w, errMsg, http.StatusBadGateway)
			return
		}

		// Calculate the backoff time (with jitter) for the next retry
		backoffTime := backoffFactor * time.Duration(1<<attempt) // Exponential backoff
		jitter := time.Duration(rand.Int63n(int64(backoffTime))) // Add some random jitter
		time.Sleep(backoffTime + jitter)
	}

	// If we reach here, it means the request was successful (or non-retryable errors were encountered)
	if resp != nil {
		defer resp.Body.Close()

		if(resp.StatusCode != 200) {
			logger.Warn.Printf("request to storage NOT successful %d", resp.StatusCode)
		}

		// Copy response headers and body from target server to client
		for key, values := range resp.Header {
			for _, v := range values {
				w.Header().Add(key, v)
			}
		}
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	}
}

func main() {
	logger.Init()

	http.HandleFunc("/upload", uploadHandler)
	logger.Info.Println("Upload Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
