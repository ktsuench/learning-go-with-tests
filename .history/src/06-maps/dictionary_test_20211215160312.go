package main

import "testing"

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("unknown")

		assertError(t, err, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {
	word := "test"
	defintion := "this is just a test"

	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		err := dict.Add(word, defintion)

		assertError(t, err, nil)
		assertDefinition(t, dict, word, defintion)
	})

	t.Run("existing word", func(t *testing.T) {
		dict := Dictionary{word: defintion}
		err := dict.Add(word, defintion)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dict, word, defintion)
	})
}

func TestUpdate(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	newDefinition := "this is just an update"

	t.Run("existing word", func(t *testing.T) {
		dict := Dictionary{word: definition}

		err := dict.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dict, word, newDefinition)
	})

	t.Run("non-existing word", func(t *testing.T) {
		dict := Dictionary{}

		err := dict.Update(word, newDefinition)

		assertError(t, err, ErrUpdateFailed)
		// assertDefinition(t, dict, word, "")
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dict Dictionary, word, definition string) {
	t.Helper()

	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("should find added word: ", err)
	}

	if got != definition {
		t.Errorf("got %q want %q", got, definition)
	}
}
