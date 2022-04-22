// Copyright (C) 2017 ScyllaDB

package jsonutil

import (
	"strings"
	"testing"
)

const sample = `
{
	"name": "Kevin",
	"cats": [
		"Belle",
		"Pedro",
		"Cassidy",
		"dogs"
	],
	"dogs": [
		"Chloe",
		"Ellie",
		"Izzy",
		"Jessie",
		"Luna"
	],
	"fish": [
		"Billy"
	]
}
`

type Person struct {
	Name string
	Dogs []string
	Cats []string
}

func TestSeekAndDecode(t *testing.T) {
	r := strings.NewReader(sample)
	decoder := NewDecoder(r)

	if err := decoder.Seek("dogs"); err != nil {
		t.Fatal(err)
	}

	var dogs []string
	for decoder.More() {
		var dog string
		if err := decoder.Decode(&dog); err != nil {
			t.Fatal(err)
		}

		dogs = append(dogs, dog)
	}

	if dogs[0] != "Chloe" {
		t.Fatalf("Expected Chloe, got %s", dogs[0])
	}

	if len(dogs) != 5 {
		t.Fatalf("Expected 5, got %d", len(dogs))
	}
}

func TestSeekNotFound(t *testing.T) {
	f := strings.NewReader(sample)
	decoder := NewDecoder(f)

	err := decoder.Seek("birds")

	if err != ErrKeyNotFound {
		t.Fatalf("expected ErrKeyNotFound, got %s", err)
	}
}
