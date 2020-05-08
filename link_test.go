package main

import "testing"

func TestCreateLinkFromFile(t *testing.T) {
	l := createLinkFromTheFile("test.csv")
	if len(l) != 96 {
		t.Errorf("Expected length of 96, but got %v", len(l))
	}
	if l[0] != "https://golang.org" {
		t.Errorf("Expected 'https://golang.org', but got %v", l[0])
	}
}
