package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	// Path to the file you want to upload
	filePath := "/Users/lbaron/Downloads/Invoice_2024-00217.pdf"
	// Field name for form-data
	fieldName := "file"
	// Target upload URL
	uploadURL := "http://localhost:8080/upload"

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a buffer and a multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create a form file field
	part, err := writer.CreateFormFile(fieldName, file.Name())
	if err != nil {
		fmt.Println("Error creating form file:", err)
		return
	}

	// Copy the file into the form file field
	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println("Error copying file:", err)
		return
	}

	// Close the multipart writer to set the terminating boundary
	err = writer.Close()
	if err != nil {
		fmt.Println("Error closing writer:", err)
		return
	}

	// Create a new POST request
	req, err := http.NewRequest("POST", uploadURL, body)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Print the response
	respBody, _ := io.ReadAll(resp.Body)
	fmt.Printf("Response: %s\n", respBody)
}
