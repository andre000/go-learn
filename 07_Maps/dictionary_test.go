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

		assertNoError(t, err)
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

	t.Run("should be able to update an existing word", func(t *testing.T) {
		err := dictionary.Update(word, newDefinition)
		assertNoError(t, err)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("should throw an error when updating a word that doesn't exist", func(t *testing.T) {
		err := dictionary.Update("unknown", newDefinition)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	dictionary := Dictionary{word: definition}

	t.Run("should be able to remove an existing word", func(t *testing.T) {
		dictionary.Delete(word)
		assertNotFound(t, dictionary, word)
	})
}

/*
	UTILS
*/

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

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("❌ received error, expected none")
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, key string, expected string) {
	t.Helper()
	received, err := dictionary.Search(key)

	assertNoError(t, err)
	assertString(t, received, expected)
}

func assertNotFound(t *testing.T, dictionary Dictionary, key string) {
	t.Helper()
	_, err := dictionary.Search(key)
	if err != ErrNotFound {
		t.Errorf("❌ word %q has been found in the dictionary", key)
	}
}
