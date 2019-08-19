package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "This is a test"}

	t.Run("should be able to search and return a word on the dictionary", func(t *testing.T) {
		received, _ := dictionary.Search("test")
		expected := "This is a test"

		assertString(t, received, expected)
	})

	t.Run("should notify the user when the word isn't found", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		expected := ErrNotFound

		assertError(t, err, expected)
	})
}

func TestAdd(t *testing.T) {
	t.Run("should be able to add a new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("test", "this is just a test")

		assertNotError(t, err)
		assertDefinition(t, dictionary, "test", "this is just a test")
	})

	t.Run("should throw an error when trying to overwrite an word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, definition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	dictionary := Dictionary{word: definition}
	newDefinition := "new definition"

	dictionary.Update(word, newDefinition)
	assertDefinition(t, dictionary, word, newDefinition)
}

func assertString(t *testing.T, received string, expected string) {
	t.Helper()
	if received != expected {
		t.Errorf("❌ received %q expected %q", received, expected)
	}
}

func assertError(t *testing.T, err error, expected error) {
	t.Helper()
	if err == nil {
		t.Fatal("❌ expected error, received none")
	}

	if err != expected {
		t.Errorf("❌ received %s expected %s", err, expected)
	}
}

func assertNotError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("❌ received error, expected none")
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, key string, expected string) {
	t.Helper()
	received, err := dictionary.Search(key)

	assertNotError(t, err)
	assertString(t, received, expected)
}
