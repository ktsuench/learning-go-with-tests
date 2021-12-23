package main

import "testing"

type StubStore struct {
	response string
}

func (s *StubStore) Fetch() string {
	return s.response
}

func TestStore(t *testing.T) {

}
