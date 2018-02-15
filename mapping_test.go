package main

import (
	"testing"
)

func CreateItem(t *testing.T) {
	s := newTestStore()
	defer s.Close()

	// Insert an item
	_, err := s.CreateMapping("http://test.com", "test", false)
	if err != nil {
		t.Fatal(err)
	}
}
