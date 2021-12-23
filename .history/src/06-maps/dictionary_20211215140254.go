package main

import (
	"errors"
)

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(key string) (string, error) {
	if val, ok := d[key]; ok {
		return val, nil
	}

	return "", ErrNotFound
}

func (d Dictionary) Add(key, val string) {

}
