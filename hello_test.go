package main

import (
	"testing"
)

func TestMessage(t *testing.T) {
	if getMessage() != "hello" {
		t.Errorf("Expected hello")
	}
}
