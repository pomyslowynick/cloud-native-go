package cmd

import "testing"

func TestPutHappyPath(t *testing.T) {
	Put("hello", "world")

	if value, _ := Get("hello"); value != "world" {
		t.Error(`Put("hello", "world"); Get("hello") != "world" `)
	}
}

func TestPutEmptyKeyAndValue(t *testing.T) {
	Put("", "")

	if value, _ := Get(""); value != "" {
		t.Error(`Put("hello", "world"); Get("hello") != "world" `)
	}
}

func TestDelete(t *testing.T) {
	Put("hello", "world")

	err := Delete("hello")
	if err != nil {
		t.Error(`Couldn't delete key "hello"`)
	}

	if value, _ := Get("hello"); value != "" {
		t.Error(`Get didn't return an empty string`)
	}
}
