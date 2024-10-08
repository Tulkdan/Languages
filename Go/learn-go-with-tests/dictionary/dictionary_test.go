package dictionary

import "testing"

func TestSearch(t *testing.T) {
    dictionary := Dictionary{"test": "this is just a test"}

    t.Run("known word", func(t *testing.T) {
        got, _ := dictionary.Search("test")
        want := "this is just a test"

        if got != want {
            t.Errorf("got %q want %q given, %q", got, want, "test")
        }
    })

    t.Run("unknown word", func(t *testing.T) {
        _, err := dictionary.Search("unknown")
        want := "could not find the word you were looking for"

        if err == nil {
            t.Fatal("expected to get an error")
        }

        if err.Error() != want {
            t.Errorf("got %q want %q given, %q", err.Error(), want, "test")
        }
    })
}

func TestAdd(t *testing.T) {
    t.Run("new word", func(t *testing.T) {
        dictionary := Dictionary{}
        dictionary.Add("test", "this is just a test")

        want := "this is just a test"
        got, err := dictionary.Search("test")
        if err != nil {
            t.Fatal("should find added word:", err)
        }

        if got != want {
            t.Errorf("got %q want %q given, %q", got, want, "test")
        }
    })

    t.Run("existing word", func(t *testing.T) {
        key := "test"
        definition := "this is just a test"
        dictionary := Dictionary{key: definition}
        err := dictionary.Add(key, "new test")

        if err != ErrWordExists {
            t.Errorf("got %q want %q", err, "")
        }

        got, _ := dictionary.Search(key)
        if got != definition {
            t.Errorf("got %q want %q given, %q", got, definition, key)
        }
    })
}
