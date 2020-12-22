package main

import (
	"io/ioutil"
	"os"
)

// Writer interface
type Writer interface {
	WriteFile(filename string, data []byte, perm os.FileMode) error
}

// FileWriter implements an abstraction of ioutil.WriterFile
type FileWriter struct {
}

// WriteFile implements the Writer interface that's been created so that ioutil.WriteFile can be mocked
func (w FileWriter) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(filename, data, perm)
}

// Reader interface
type Reader interface {
	ReadFile(filename string) ([]byte, error)
}

// FileReader implements an abstraction of ioutil.ReaderFile
type FileReader struct {
}

// ReadFile implements the Reader interface that's been created so that ioutil.ReadFile can be mocked
func (w FileReader) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}
