package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
	first := d[0]
	if first != "Ace of Spades" {
		t.Errorf("Expected first card Ace of Spades, but got %v", first)
	}
	last := d[len(d)-1]
	if last != "Four of Clubs" {
		t.Errorf("Expected first card Four of Clubs, but got %v", last)
	}
}

// source: https://willdot.net/29MockingAWriterInGo/
type FakeFileWriter struct {
}

func (f FakeFileWriter) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return nil
}

func TestSaveToFile(t *testing.T) {
	d := newDeck()
	r := d.saveToFile(FakeFileWriter{}, "test")
	fmt.Println(r)
}

type FakeFileReader struct {
}

func (f FakeFileReader) ReadFile(filename string) ([]byte, error) {
	d := newDeck()
	return []byte(d.toString()), nil
}
func TestLoadDeckFromFile(t *testing.T) {
	d := loadDeckFromFile(FakeFileReader{}, "test")
	fmt.Println(d)
}
