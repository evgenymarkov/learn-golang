package dictionary

import "testing"

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}

		word := "test"
		definition := "this is just test"

		err := dictionary.Add(word, definition)

		assertNoError(t, err)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just test"}

		word := "test"
		definition := "this is another test definition"

		err := dictionary.Add(word, definition)

		assertError(t, err, ErrWordExist)
	})
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just test"}

	t.Run("known word", func(t *testing.T) {
		word := "test"
		definition := "this is just test"

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("unknown word", func(t *testing.T) {
		word := "unknown"
		_, err := dictionary.Search(word)

		assertError(t, err, ErrNoWordFound)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		word := "test"
		definition := "this is just test"
		dictionary := Dictionary{word: definition}

		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)

		assertNoError(t, err)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("unknown word", func(t *testing.T) {
		word := "test"
		definition := "this is just test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)

	assertError(t, err, ErrNoWordFound)
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Error("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word string, want string) {
	t.Helper()

	got, err := dictionary.Search(word)

	assertNoError(t, err)

	if got != want {
		t.Errorf("got %q, want %q, given %q", got, want, "test")
	}
}
