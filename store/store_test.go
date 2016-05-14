package store_test

import (
	"github.com/ethanaubuchon/lightdm/store"
	"testing"
)

func TestGet(t *testing.T) {
	s := store.NewStore()
	actual, err := s.Get("A")
	expected := ""

	if err == nil {
		t.Error("Expected an error")
	}

	if actual != expected {
		t.Error("Expected " + expected + " Actual " + actual)
	}
}

func TestSet(t *testing.T) {
	s := store.NewStore()
	expected := "value"
	key := "Key"
	s.Set(key, expected)
	actual, err := s.Get(key)

	if err != nil {
		t.Error("Unexpected Error")
	}

	if actual != expected {
		t.Error("Expected " + expected + " Actual " + actual)
	}
}

func TestUnset(t *testing.T) {
	s := store.NewStore()
	expected := ""
	key := "Key"
	s.Set(key, "value")
	s.Unset(key)
	actual, err := s.Get(key)

	if err == nil {
		t.Error("Expected an error")
	}

	if actual != expected {
		t.Error("Expected " + expected + " Actual " + actual)
	}
}
