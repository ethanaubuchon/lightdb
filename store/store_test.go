package store_test

import (
	"github.com/ethanaubuchon/lightdm/store"
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	s := store.NewStore()
	actual, err := s.Get("A")
	expected := ""

	if err == nil {
		t.Error("Expected an error")
	}

	if actual.(string) != expected {
		t.Error("Expected ", expected, " Actual ", actual)
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

	if actual.(string) != expected {
		t.Error("Expected ", expected, " Actual ", actual)
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

	if actual.(string) != expected {
		t.Error("Expected ", expected, " Actual ", actual)
	}
}

func TestTransactionIsolation(t *testing.T) {
	s := store.NewStore()
	s.Set("test", "ABCD")
	s.Set("other", 1234)
	s.Set("thisOne", "Unicorn")
	s.Set("WhatsThis", true)

	tx := s.Begin()
	tx.Set("foo", "bar")
	tx.Set("bar", "foo")
	tx.Set("test", false)
	tx.Set("other", "ABC News at 6")
	tx.Unset("other")
	tx.Unset("bar")
	tx.Unset("WhatsThis")

	actual, err := s.Get("test")
	if err != nil {
		t.Error("Unexpected Error")
	}

	var expected interface{} = "ABCD"
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Expected ", expected, " Actual ", actual)
	}

	actual, err = s.Get("other")
	if err != nil {
		t.Error("Unexpected Error")
	}

	expected = 1234
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Expected ", expected, " Actual ", actual)
	}

	actual, err = s.Get("thisOne")
	if err != nil {
		t.Error("Unexpected Error")
	}

	expected = "Unicorn"
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Expected ", expected, " Actual ", actual)
	}

	actual, err = s.Get("WhatsThis")
	if err != nil {
		t.Error("Unexpected Error")
	}

	expected = true
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Expected ", expected, " Actual ", actual)
	}

	actual, err = tx.Get("test")
	if err != nil {
		t.Error("Unexpected Error")
	}

	expected = false
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Expected ", expected, " Actual ", actual)
	}

	actual, err = tx.Get("other")
	if err != nil {
		t.Error("Unexpected Error")
	}

	expected = nil
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Expected ", expected, " Actual ", actual)
	}

	actual, err = tx.Get("thisOne")
	if err != nil {
		t.Error("Unexpected Error")
	}

	expected = "Unicorn"
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Expected ", expected, " Actual ", actual)
	}

	actual, err = tx.Get("WhatsThis")
	if err != nil {
		t.Error("Unexpected Error")
	}

	expected = nil
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Expected ", expected, " Actual ", actual)
	}

	actual, err = tx.Get("foo")
	if err != nil {
		t.Error("Unexpected Error")
	}

	expected = "bar"
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Expected ", expected, " Actual ", actual)
	}

	actual, err = tx.Get("bar")
	if err != nil {
		t.Error("Unexpected Error")
	}

	expected = nil
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Expected ", expected, " Actual ", actual)
	}

}

func TestRollback(t *testing.T) {
	s := store.NewStore()
	s.Set("test", "ABCD")
	s.Set("other", 1234)
	s.Set("thisOne", "Unicorn")
	s.Set("WhatsThis", true)

	tx := s.Begin()
	tx.Set("foo", "bar")
	tx.Set("bar", "foo")
	tx.Set("test", false)
	tx.Set("other", "ABC News at 6")
	tx.Unset("other")
	tx.Unset("bar")
	tx.Unset("WhatsThis")

	tx.Rollback()

	actual, err := s.Get("test")
	if err != nil {
		t.Error("Unexpected Error")
	}

	var expected interface{} = "ABCD"
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Expected ", expected, " Actual ", actual)
	}

	actual, err = s.Get("other")
	if err != nil {
		t.Error("Unexpected Error")
	}

	expected = 1234
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Expected ", expected, " Actual ", actual)
	}

	actual, err = s.Get("thisOne")
	if err != nil {
		t.Error("Unexpected Error")
	}

	expected = "Unicorn"
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Expected ", expected, " Actual ", actual)
	}

	actual, err = s.Get("WhatsThis")
	if err != nil {
		t.Error("Unexpected Error")
	}

	expected = true
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Expected ", expected, " Actual ", actual)
	}
}

func TestCommit(t *testing.T) {
	s := store.NewStore()
	s.Set("test", "ABCD")
	s.Set("other", 1234)
	s.Set("thisOne", "Unicorn")
	s.Set("WhatsThis", true)

	tx := s.Begin()
	tx.Set("foo", "bar")
	tx.Set("bar", "foo")
	tx.Set("test", false)
	tx.Set("other", "ABC News at 6")
	tx.Unset("other")
	tx.Unset("bar")
	tx.Unset("WhatsThis")

	tx.Commit()

	actual, err := s.Get("test")
	if err != nil {
		t.Error("Unexpected Error")
	}

	var expected interface{} = false
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Expected ", expected, " Actual ", actual)
	}

	actual, err = s.Get("other")
	if err != nil {
		t.Error("Unexpected Error")
	}

	expected = nil
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Expected ", expected, " Actual ", actual)
	}

	actual, err = s.Get("thisOne")
	if err != nil {
		t.Error("Unexpected Error")
	}

	expected = "Unicorn"
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Expected ", expected, " Actual ", actual)
	}

	actual, err = s.Get("WhatsThis")
	if err != nil {
		t.Error("Unexpected Error")
	}

	expected = nil
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Expected ", expected, " Actual ", actual)
	}

	actual, err = s.Get("foo")
	if err != nil {
		t.Error("Unexpected Error")
	}

	expected = "bar"
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Expected ", expected, " Actual ", actual)
	}

	actual, err = s.Get("bar")
	if err != nil {
		t.Error("Unexpected Error")
	}

	expected = nil
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Expected ", expected, " Actual ", actual)
	}

}
