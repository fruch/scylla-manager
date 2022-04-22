// Copyright (C) 2017 ScyllaDB

package jsonutil

import (
	"encoding/json"
	"errors"
	"io"
)

// ErrKeyNotFound is an error if a key does not exist.
var ErrKeyNotFound = errors.New("key not found")

// Decoder is a wrapper around a json.Decoder.
type Decoder struct {
	*json.Decoder
}

// NewDecoder creates a new JSON decoder.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{json.NewDecoder(r)}
}

// Seek seeks forward-only to a key in the JSON.
func (d *Decoder) Seek(key string) error {
	for t, err := d.Token(); err == nil; t, err = d.Token() {
		if s, ok := t.(string); ok && s == key {
			t, err = d.Token()
			if err != nil {
				return err
			}

			if x, ok := t.(json.Delim); ok && (x == '{' || x == '[') {
				return nil
			}
		}
	}

	return ErrKeyNotFound
}
