package main

import (
	"errors"
)

type Dictionary map[string]string

var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

func (d Dictionary) Search(word string) (string, error) {
	if val, ok := d[word]; ok {
		return val, nil
	}

	return "", ErrNotFound
}

func (d Dictionary) Add(word, definition string) error {
	if _, exists := d[word]; exists {
		return ErrWordExists
	}

	d[word] = definition
	return nil
}
