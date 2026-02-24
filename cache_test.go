package main

import (
	"slices"
	"testing"
)

func TestWriteFile(t *testing.T) {
	tweets := []Tweet{
		{Url: "test1"},
		{Url: "test2"},
	}

	err := WriteCache(CACHE_FILE, tweets)

	if err != nil {
		t.Fatal(err)
	}
}

func TestReadFile(t *testing.T) {
	tweets := []Tweet{
		{Url: "test1"},
		{Url: "test2"},
	}

	result := GetCache(CACHE_FILE)

	t.Logf("result: %v", result)

	if !slices.Equal(result, tweets) {
		t.Fatal("result were not equal")
	}
}
