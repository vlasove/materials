package dictionary

import "testing"

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, err := dictionary.Search("test")
		want := "this is just a test"
		assertError(t, err, nil)
		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("bob")

		assertError(t, err, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		old := "this is just a test"
		dictionary := Dictionary{word: old}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, old)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("update existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new test"

		err := dictionary.Update(word, newDefinition)
		assertError(t, err, nil)
		got, err := dictionary.Search(word)

		assertError(t, err, nil)
		assertString(t, got, newDefinition)
	})

	t.Run("update not existing word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordNotExists)

		_, err = dictionary.Search(word)
		assertError(t, err, ErrNotFound)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Delete(word)
		assertError(t, err, nil)
		_, err = dictionary.Search(word)

		assertError(t, err, ErrNotFound)

	})

	t.Run("delete not existing word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		err := dictionary.Delete(word)
		assertError(t, err, ErrWordNotExists)

		_, err = dictionary.Search(word)
		assertError(t, err, ErrNotFound)
	})
}
