package main

import (
	"testing"
)

func TestAddr(t *testing.T) {
	var expected int = 5
	var actual int = tambah(2, 3)

	if expected != actual {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
