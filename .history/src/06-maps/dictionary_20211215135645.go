package main

import (
	"errors"
)

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	if val, ok := d[key]; ok {
		return val, nil
	}

	return "", errors.New("could not find the word you were looking for")
}
