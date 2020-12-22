package main

import "testing"

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
