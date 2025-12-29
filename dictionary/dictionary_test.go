package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error.")
		}

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
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})

}
func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})

}
func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "test definition"}

		err := dictionary.Delete(word)

		assertError(t, err, nil)

		_, err = dictionary.Search(word)

		assertError(t, err, ErrNotFound)
	})

	t.Run("non-existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}

		err := dictionary.Delete(word)

		assertError(t, err, ErrWordDoesNotExist)
	})
}
func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot perform operation on word because it does not exist")
)

func (d Dictionary) Search(key string) (string, error) {
	def, ok := d[key]
	if ok {
		return def, nil
	} else {
		return "", ErrNotFound
	}
}
func (d Dictionary) Add(key, data string) error {
	switch _, err := d.Search(key); err {
	case ErrNotFound:
		d[key] = data
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}
func (d Dictionary) Update(key, data string) error {
	switch _, err := d.Search(key); err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[key] = data
	default:
		return err
	}
	return nil
}
func (d Dictionary) Delete(key string) error {
	switch _, err := d.Search(key); err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, key)
	default:
		return err
	}
	return nil
}

/*
An interesting property of maps is that you can modify them without passing as an address to it (e.g &myMap)
This may make them feel like a "reference type",
but as Dave Cheney describes they are not.
A map value is a pointer to a runtime.hmap structure.
So when you pass a map to a function/method, you are indeed copying it,
 but just the pointer part, not the underlying data structure that contains the data.

A gotcha with maps is that they can be a nil value.
A nil map behaves like an empty map when reading,
but attempts to write to a nil map will cause a runtime panic.

Therefore, you should never initialize a nil map variable:

var m map[string]string

Instead, you can initialize an empty map or use the make keyword to create a map for you:

var dictionary = map[string]string{}
// OR
var dictionary = make(map[string]string)
*/
