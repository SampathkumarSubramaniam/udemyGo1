package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	testDeck := newDeck()
	if len(testDeck) != 12 {
		t.Error("Expected deck length of 16, but got", len(testDeck))
	}
	if testDeck[0] != "Ace-Diamond" {
		t.Error("Expected Ace-Diamond, but got", testDeck[0])
	}
	if testDeck[len(testDeck)-1] != "Three-someThingElse" {
		t.Error("Expected Ace-Diamond, but got", testDeck[0])
	}
}
